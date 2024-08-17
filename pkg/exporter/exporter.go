package exporter

import (
	"go.uber.org/zap"
)

type StockPriceExporter struct {
	Logger *zap.Logger
	Prices []float64
	WatchList []string
}

func NewStockPriceExporter(watchList []string) *StockPriceExporter {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// Create a new exporter
	exporter := StockPriceExporter{
		Logger: logger,
		Prices: make([]float64, len(watchList)),
		WatchList: watchList,
	}

	return &exporter
}
