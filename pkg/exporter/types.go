package exporter

type YFinanceChart struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				FullExchangeName     string  `json:"fullExchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				HasPrePostMarketData bool    `json:"hasPrePostMarketData"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				FiftyTwoWeekHigh     float64 `json:"fiftyTwoWeekHigh"`
				FiftyTwoWeekLow      float64 `json:"fiftyTwoWeekLow"`
				RegularMarketDayHigh float64 `json:"regularMarketDayHigh"`
				RegularMarketDayLow  float64 `json:"regularMarketDayLow"`
				RegularMarketVolume  int     `json:"regularMarketVolume"`
				LongName             string  `json:"longName"`
				ShortName            string  `json:"shortName"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PreviousClose        float64 `json:"previousClose"`
				Scale                int     `json:"scale"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
				TradingPeriods [][]struct {
					Timezone  string `json:"timezone"`
					Start     int    `json:"start"`
					End       int    `json:"end"`
					Gmtoffset int    `json:"gmtoffset"`
				} `json:"tradingPeriods"`
				DataGranularity string   `json:"dataGranularity"`
				Range           string   `json:"range"`
				ValidRanges     []string `json:"validRanges"`
			} `json:"meta"`
			Timestamp  []int `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					High   []any `json:"high"`
					Low    []any `json:"low"`
					Open   []any `json:"open"`
					Volume []any `json:"volume"`
					Close  []any `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
		Error any `json:"error"`
	} `json:"chart"`
}

type YFinanceQuote struct {
	Explains []any `json:"explains"`
	Count    int   `json:"count"`
	Quotes   []struct {
		Exchange       string  `json:"exchange"`
		Shortname      string  `json:"shortname"`
		QuoteType      string  `json:"quoteType"`
		Symbol         string  `json:"symbol"`
		Index          string  `json:"index"`
		Score          float64 `json:"score"`
		TypeDisp       string  `json:"typeDisp"`
		Longname       string  `json:"longname"`
		ExchDisp       string  `json:"exchDisp"`
		IsYahooFinance bool    `json:"isYahooFinance"`
	} `json:"quotes"`
	News []struct {
		UUID                string `json:"uuid"`
		Title               string `json:"title"`
		Publisher           string `json:"publisher"`
		Link                string `json:"link"`
		ProviderPublishTime int    `json:"providerPublishTime"`
		Type                string `json:"type"`
		Thumbnail           struct {
			Resolutions []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
				Tag    string `json:"tag"`
			} `json:"resolutions"`
		} `json:"thumbnail"`
		RelatedTickers []string `json:"relatedTickers"`
	} `json:"news"`
	Nav                            []any `json:"nav"`
	Lists                          []any `json:"lists"`
	ResearchReports                []any `json:"researchReports"`
	ScreenerFieldResults           []any `json:"screenerFieldResults"`
	TotalTime                      int   `json:"totalTime"`
	TimeTakenForQuotes             int   `json:"timeTakenForQuotes"`
	TimeTakenForNews               int   `json:"timeTakenForNews"`
	TimeTakenForAlgowatchlist      int   `json:"timeTakenForAlgowatchlist"`
	TimeTakenForPredefinedScreener int   `json:"timeTakenForPredefinedScreener"`
	TimeTakenForCrunchbase         int   `json:"timeTakenForCrunchbase"`
	TimeTakenForNav                int   `json:"timeTakenForNav"`
	TimeTakenForResearchReports    int   `json:"timeTakenForResearchReports"`
	TimeTakenForScreenerField      int   `json:"timeTakenForScreenerField"`
	TimeTakenForCulturalAssets     int   `json:"timeTakenForCulturalAssets"`
	TimeTakenForSearchLists        int   `json:"timeTakenForSearchLists"`
}
