package exporter

import (
	"fmt"
	"net/http"
	"strings"
)

// SymbolHandler returns prices of symbols in our WatchList
func (s *StockPriceExporter) SymbolHandler(w http.ResponseWriter, r *http.Request) {
	var output []string
	for i, symbol := range(s.WatchList) {
		output = append(output, fmt.Sprintf("prices{symbol=\"%s\"} %v\n", symbol, s.Prices[i]))
	}
	
	w.Write([]byte(strings.Join(output, "")))
}

// Serve uses prometheus's prom-auto package to serve metrics.
// It is a wrapper around the stock-exporter.
func (s *StockPriceExporter) Serve(address, port string) {
	go s.updateStockPrice()

	s.Logger.Sugar().Infof("StockPriceExporter listening on %s:%s", address, port)

	http.HandleFunc("/metrics", s.SymbolHandler)

	s.Logger.Sugar().Fatal(http.ListenAndServe(address+":"+port, nil))
}