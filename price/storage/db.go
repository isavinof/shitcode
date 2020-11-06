package storage

type DB struct {
}

var DBInstance = &DB{}

func (d *DB) Set(symbol string, price *Price) error {
	_, err := MySQLInstance.Query("insert bla bla %v %v", symbol, price)
	return err
}

func (d *DB) Get(symbol string) (*Price, error) {
	bidAsk := 0.0
	price := Price{Ask: &bidAsk, Bid: &bidAsk}
	rows, err := MySQLInstance.Query("select * from ...", symbol)
	if err != nil {
		return nil, err
	}

	err = rows.Scan(price.Bid, price.Ask)
	return &price, err
}
