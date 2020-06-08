package asset

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StockMaster struct {
	Id              int64  `json:"id"`
	ExchangeCountry string `json:"exchangeCountry"`
	Timestamp       int64  `json:"timestamp"`
	Count           int64  `json:"count"`
	SpotCount       int64  `json:"spotCount"`
	IndexCount      int64  `json:"indexCount"`
}

type Asset struct {
	ExchangeCountry      string  `json:"exchangeCountry"`
	SecurityType         string  `json:"securityType"`
	Market               string  `json:"market"`
	TickerSymbol         string  `json:"tickerSymbol"`
	ShortCode            string  `json:"shortCode"`
	KoreanName           string  `json:"koreanName"`
	EnglishName          string  `json:"englishName,omitempty"`
	CountryCode          string  `json:"countryCode"`
	CurrencyISOCode      string  `json:"currencyISOCode"`
	ListedSharesCount    int64   `json:"listedSharesCount"`
	BasePrice            float64 `json:"basePrice"`
	IsTradingSuspended   bool    `json:"isTradingSuspended"`
	SalesDate            string  `json:"salesDate"`
	SectorCode           string  `json:"sectorCode"`
	IsManufacturingOrSme bool    `json:"isManufacturingOrSme"`
	PreviousClosingPrice float64 `json:"previousClosingPrice"`
	DecimalPlace         int64   `json:"decimalPlace"`
	IsIndex              bool    `json:"isIndex"`
	Timestamp            int64   `json:"timestamp"`
	Code                 string  `json:"code"`
	CommonStockCode      string  `json:"commonStockCode"`
	IsCommonStock        bool    `json:"isCommonStock"`
	NewListing           bool    `json:"newListing"`
	Delisted             bool    `json:"delisted"`
}

func FetchStockMasters() (*[]StockMaster, error) {
	resp, err := http.Get("https://quotation-static.dunamu.com/stockmaster_timestamp")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []StockMaster
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *StockMaster) FetchAssets() (*[]Asset, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://quotation-static.dunamu.com/%s_stockmaster", m.ExchangeCountry),
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []Asset
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
