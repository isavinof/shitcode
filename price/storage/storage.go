package storage

import "strings"

type Storage struct {
}

var StorageInstance = Storage{}

func (s *Storage) Set(symbol string, price *Price) error {
	err := CacheInstance.Set(symbol, price)
	if err != nil {
		return err
	}

	return DBInstance.Set(symbol, price)
}

func (s *Storage) Get(symbol string) (*Price, error) {
	price, err := CacheInstance.Get(symbol)
	if err == nil {
		return price, nil
	}

	if !strings.Contains(err.Error(), "redis: nil") {
		return nil, err
	}

	return DBInstance.Get(symbol)
}

func (s *Storage) MGet(symbols []string) ([]*Price, error) {
	var res []*Price
	for _, symbol := range symbols {
		p, err := s.Get(symbol)
		if err != nil {
			return nil, err
		}
		res = append(res, p)
	}

	return res, nil
}
