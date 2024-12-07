package exporter

import (
	"fmt"
	"net/http"
	"strings"
)

// Healthz returns status ok to satisfy health checks
func (s *StockPriceExporter) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

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
	http.HandleFunc("/healthz", s.Healthz)

	s.Logger.Sugar().Fatal(http.ListenAndServe(address+":"+port, nil))
}