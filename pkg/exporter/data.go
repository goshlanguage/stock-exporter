package exporter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	marketOpen = int(14.5 * 60.0 * 60.0) // 14:30 UTC in seconds
	marketClose = 20 * 60 * 60 // 20:00 UTC in seconds
)

// updateStockPrice updates the stock price for each symbol in our WatchList and it's associated gauge
func (s *StockPriceExporter) updateStockPrice() {
	// pollPeriod is adjusted per the market open/close time
	// It defaults to 1 minute
	// Switches to 30 minutes outside of the marketOpen and marketClose
	pollPeriod := 1 * time.Minute

	for {
		// checking time, only run once every 30 minutes when the market is closed
		now := time.Now()

		startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
    secondsElapsed := int(now.Sub(startOfToday).Seconds())

		dayOfWeek := now.Weekday().String()
		s.Logger.Sugar().Infof("Today is %s", dayOfWeek)
    
		if (secondsElapsed > marketOpen && secondsElapsed < marketClose) &&
			(dayOfWeek != "Sunday" && dayOfWeek != "Saturday") {
			if pollPeriod != time.Minute {
				s.Logger.Sugar().Infof("Market open detected, updating poll period to 1 minute")
			}

			pollPeriod = 1 * time.Minute
		} else {
			if pollPeriod != time.Minute {
				s.Logger.Sugar().Infof("Market close detected, updating poll period to 30 minutes")
			}
		
			pollPeriod = 30 * time.Minute
		}
		
		s.Logger.Sugar().Infof("Updating stock prices for %d symbols", len(s.WatchList))

		for i, symbol := range s.WatchList {
			// url := "https://query2.finance.yahoo.com/v1/finance/search?q=" + symbol
			url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s", symbol)

			resp, err := http.Get(url)
			statusCode := resp.StatusCode
			if err != nil || statusCode != 200 {
				s.Logger.Sugar().Errorf("Failed to lookup symbol %s: %s\tStatus: %v\tURL: %s\n", symbol, err, statusCode,url)
				continue
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
					s.Logger.Sugar().Errorf("Failed to read body for symbol %s:  %s\n", symbol, err)
					continue
			}

			var chart *YFinanceChart
			
			err = json.Unmarshal(body, &chart)
			if err != nil {
					s.Logger.Sugar().Errorf("Failed to unmarshal symbol %s: %s\n", symbol, err)
					continue
			}

			s.Logger.Sugar().Infow("Current price", "Symbol", symbol, "price", chart.Chart.Result[0].Meta.RegularMarketPrice)
	
			s.Prices[i] = chart.Chart.Result[0].Meta.RegularMarketPrice

			resp.Body.Close()
		}

		time.Sleep(pollPeriod)
	}
}
