package asset

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Security struct {
	Name                 string  `json:"name"`
	Code                 string  `json:"code"`
	TickerSymbol         string  `json:"tickerSymbol"`
	ShortCode            string  `json:"shortCode"`
	Date                 string  `json:"date"`
	IsIndex              bool    `json:"isIndex"`
	ExchangeCountry      string  `json:"exchangeCountry"`
	SecurityType         string  `json:"securityType"`
	Market               string  `json:"market"`
	GBoard               string  `json:"gBoard"`
	TradingHours         string  `json:"tradingHours"`
	OpeningPrice         float64 `json:"openingPrice"`
	HighPrice            float64 `json:"highPrice"`
	LowPrice             float64 `json:"lowPrice"`
	TradePrice           float64 `json:"tradePrice"`
	TradeVolume          int64   `json:"tradeVolume"`
	AccTradeVolume       int64   `json:"accTradeVolume"`
	AccTradePrice        float64 `json:"accTradePrice"`
	GlobalAccTradeVolume int64   `json:"globalAccTradeVolume"`
	GlobalAccTradePrice  float64 `json:"globalAccTradePrice"`
	ExpectedTradePrice   float64 `json:"expectedTradePrice"`
	ExpectedTradeVolume  int64   `json:"expectedTradeVolume"`
	Change               string  `json:"change"`
	ChangePrice          float64 `json:"changePrice"`
	PrevClosingPrice     float64 `json:"prevClosingPrice"`
	AccAskVolume         int64   `json:"accAskVolume"`
	AccBidVolume         int64   `json:"accBidVolume"`
	StandardPrice        float64 `json:"standardPrice"`
	TradeTime            string  `json:"tradeTime"`
	DisplayTime          string  `json:"displayTime"`
	RegularHoursStatus   string  `json:"regularHoursStatus"`
	ListedShareCount     int64   `json:"listedShareCount"`
	PrevListedShareCount int64   `json:"prevListedShareCount"`
	High52wPrice         float64 `json:"high52wPrice"`
	High52wDate          string  `json:"high52wDate"`
	Low52wPrice          float64 `json:"low52wPrice"`
	Low52wDate           string  `json:"low52wDate"`
	IsTradingSuspended   bool    `json:"isTradingSuspended"`
	IsDelisted           bool    `json:"isDelisted"`
	DecimalPlace         int64   `json:"decimalPlace"`
	ReferredType         string  `json:"referredType"`
	Timestamp            int64   `json:"timestamp"`
	CreatedAt            string  `json:"createdAt"`
	ModifiedAt           string  `json:"modifiedAt"`
	NewListing           bool    `json:"newListing"`
	SignedChangePrice    float64 `json:"signedChangePrice"`
	SignedChangeRate     float64 `json:"signedChangeRate"`
	CommonStockCode      string  `json:"commonStockCode"`
	ExpectedChangePrice  float64 `json:"expectedChangePrice"`
	ExpectedChangeRate   float64 `json:"expectedChangeRate"`
	IsExpected           bool    `json:"isExpected"`
	MarketCap            float64 `json:"marketCap"`
	ChangeRate           float64 `json:"changeRate"`
}

const quotApi = "https://quotation-api.dunamu.com/v1/recent/securities"

func FetchSecuritiesByCodes(codes []string) (*[]Security, error) {
	resp, err := http.Get(quotApi + "?codes=" + strings.Join(codes, ","))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []Security
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}