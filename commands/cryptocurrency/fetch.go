package cryptocurrency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type MarketMaster struct {
	Market      string `json:"market"`       //"KRW-BTC"
	KoreanName  string `json:"korean_name"`  // 비트코인
	EnglishName string `json:"english_name"` // Bitcoin
}

type Ticker struct {
	Market             string  `json:"market"`
	TradeDate          string  `json:"trade_date"`
	TradeTime          string  `json:"trade_time"`
	TradeDateKst       string  `json:"trade_date_kst"`
	TradeTimeKst       string  `json:"trade_time_kst"`
	TradeTimestamp     int64   `json:"trade_timestamp"`
	OpeningPrice       float64 `json:"opening_price"`
	HighPrice          float64 `json:"high_price"`
	LowPrice           float64 `json:"low_price"`
	TradePrice         float64 `json:"trade_price"`
	PrevClosingPrice   float64 `json:"basePrice"`
	Change             string  `json:"change"`
	ChangePrice        float64 `json:"change_price"`
	ChangeRate         float64 `json:"change_rate"`
	SignedChangePrice  float64 `json:"signed_change_price"`
	SignedChangeRate   float64 `json:"signed_change_rate"`
	TradeVolume        float64 `json:"trade_volume"`
	AccTradePrice      float64 `json:"acc_trade_price"`
	AccTradePrice24h   float64 `json:"acc_trade_price_24h"`
	AccTradeVolume     float64 `json:"acc_trade_volume"`
	AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`
	Highest52WeekPrice float64 `json:"highest_52_week_price"`
	Highest52WeekDate  string  `json:"highest_52_week_date"`
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`
	Timestamp          int64   `json:"timestamp"`
}

func FetchMarketMasters() (*[]MarketMaster, error) {
	resp, err := http.Get("https://api.upbit.com/v1/market/all")
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result []MarketMaster
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func FetchTicker(markets []string) (*[]Ticker, error) {
	resp, err := http.Get(
		fmt.Sprintf("https://api.upbit.com/v1/ticker?markets=%s", strings.Join(markets, ",")),
	)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []Ticker
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
