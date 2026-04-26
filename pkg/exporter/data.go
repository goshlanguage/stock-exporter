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
	maxRetries = 3
	baseDelay = 1 * time.Second
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
			
			var resp *http.Response
			var err error
			var statusCode int
			
			// Retry logic for handling rate limiting and transient errors
			for attempt := 0; attempt <= maxRetries; attempt++ {
				resp, err = http.Get(url)
				if err != nil {
					s.Logger.Sugar().Errorf("Attempt %d: Failed to lookup symbol %s: %s\tURL: %s\n", attempt+1, symbol, err, url)
					if attempt < maxRetries {
						delay := time.Duration(attempt+1) * baseDelay
						s.Logger.Sugar().Infof("Waiting %v before retrying symbol %s\n", delay, symbol)
						time.Sleep(delay)
						continue
					} else {
						break
					}
				}
				
				statusCode = resp.StatusCode
				if statusCode == 200 {
					// Success case
					break
				} else if statusCode == 429 {
					// Rate limiting - exponential backoff
					s.Logger.Sugar().Warnf("Rate limited (429) for symbol %s, attempt %d\n", symbol, attempt+1)
					if attempt < maxRetries {
						delay := time.Duration(attempt+1) * baseDelay * 2 // Exponential backoff
						s.Logger.Sugar().Infof("Waiting %v before retrying symbol %s\n", delay, symbol)
						time.Sleep(delay)
						continue
					} else {
						s.Logger.Sugar().Errorf("Max retries reached for symbol %s after %d attempts\n", symbol, maxRetries+1)
						resp.Body.Close()
						break
					}
				} else {
					// Other HTTP error
					s.Logger.Sugar().Errorf("Failed to lookup symbol %s: Status %v\tURL: %s\n", symbol, statusCode, url)
					resp.Body.Close()
					break
				}
			}
			
			// Process successful response
			if resp != nil && resp.StatusCode == 200 {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					s.Logger.Sugar().Errorf("Failed to read body for symbol %s:  %s\n", symbol, err)
					resp.Body.Close()
					continue
				}

				var chart *YFinanceChart
				
				err = json.Unmarshal(body, &chart)
				if err != nil {
					s.Logger.Sugar().Errorf("Failed to unmarshal symbol %s: %s\n", symbol, err)
					resp.Body.Close()
					continue
				}

				s.Logger.Sugar().Infow("Current price", "Symbol", symbol, "price", chart.Chart.Result[0].Meta.RegularMarketPrice)
		
				s.Prices[i] = chart.Chart.Result[0].Meta.RegularMarketPrice

				resp.Body.Close()
			} else if resp != nil {
				// Close response body for non-200 responses
				resp.Body.Close()
			}
		}

		time.Sleep(pollPeriod)
	}
}
