package price

import "github.com/gtforge/isavinof/test/price/storage"

type Provider struct {
}

func NewProvider(c chan storage.SymbolPrice) Provider {
	return Provider{}
}

func (p *Provider) Consume() error {
	for {
		if 1 == 0 {
			return nil
		}
	}
}
