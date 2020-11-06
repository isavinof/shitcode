package shitcode

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gtforge/isavinof/test/price"
	"github.com/gtforge/isavinof/test/price/storage"
)

func main() {
	priceChan := make(chan storage.SymbolPrice, 10)

	provider := price.NewProvider(priceChan)
	go provider.Consume()

	for {
		symbolPrice, ok := <-priceChan
		if !ok {
			break
		}
		err := storage.StorageInstance.Set(symbolPrice.Symbol, &symbolPrice.Price)
		if err != nil {
			log.Panic("error to set price:", err)
		}
	}
}

func HTTPHandlerGetPrices(writer http.ResponseWriter, request *http.Request) {
	symbols := []string{} // get from request
	prices, err := storage.StorageInstance.MGet(symbols)
	if err != nil {
		writer.Write([]byte(err.Error()))
		writer.WriteHeader(500)
		return
	}

	res, err := json.Marshal(prices)
	if err != nil {
		writer.Write([]byte(err.Error()))
		writer.WriteHeader(500)
		return
	}

	writer.Write(res)
	writer.WriteHeader(200)
	return
}
