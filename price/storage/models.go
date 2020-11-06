package storage

type SymbolPrice struct {
	Price
	Symbol string
}

type Price struct {
	Bid *float64
	Ask *float64
}
