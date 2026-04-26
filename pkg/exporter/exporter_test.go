package exporter

import (
	"testing"
)

func TestStockPriceExporter(t *testing.T) {
	// This is a placeholder test - in a real scenario you would mock the HTTP calls
	// and test the retry logic and error handling
	watchList := []string{"AMZN", "MSFT"}
	exporter := NewStockPriceExporter(watchList)
	
	if exporter == nil {
		t.Error("Expected exporter to be created successfully")
	}
	
	if len(exporter.WatchList) != 2 {
		t.Errorf("Expected watch list length to be 2, got %d", len(exporter.WatchList))
	}
}