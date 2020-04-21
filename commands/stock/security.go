package stock

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type FetchedSecuritiesResult struct {
	RecentSecurities []FetchedSecurity `json:"recentSecurities"`
}

type FetchedSecurity struct {
	AccTradeVolume        int64   `json:"accTradeVolume"`
	Board                 string  `json:"board"`
	Change                string  `json:"change"`
	ChangePrice           float64 `json:"changePrice"`
	ChangePriceRate       float64 `json:"changePriceRate"`
	Code                  string  `json:"code"`
	Currency              string  `json:"currency"`
	Date                  string  `json:"date"`
	DayChartUrl           string  `json:"dayChartUrl"`
	DelayedMinutes        int     `json:"delayedMinutes"`
	DisplayedPrice        float64 `json:"displayedPrice"`
	Eps                   int     `json:"eps"`
	ExchangeCountry       string  `json:"exchangeCountry"`
	ExchangeCountryName   string  `json:"exchangeCountryName"`
	ForeignRatio          string  `json:"foreignRatio"`
	GlobalAccTradePrice   float64 `json:"globalAccTradePrice"`
	High52wPrice          float64 `json:"high52wPrice"`
	HighPrice             float64 `json:"highPrice"`
	Id                    string  `json:"id"`
	IsIndex               bool    `json:"isIndex"`
	IsVi                  bool    `json:"isVi"`
	Low52wPrice           float64 `json:"low52wPrice"`
	LowPrice              float64 `json:"lowPrice"`
	Market                string  `json:"market"`
	MarketCapRank         int     `json:"marketCapRank"`
	MarketName            string  `json:"marketName"`
	MarketWarningMsg      string  `json:"marketWarningMsg"`
	MiniDayChartUrl       string  `json:"miniDayChartUrl"`
	MiniDayGuidedChartUrl string  `json:"miniDayGuidedChartUrl"`
	Name                  string  `json:"name"`
	OpeningPrice          float64 `json:"openingPrice"`
	Per                   float64 `json:"per"`
	PrevClosingPrice      float64 `json:"prevClosingPrice"`
	RegularHoursStatus    string  `json:"regularHoursStatus"`
	SectorName            string  `json:"sectorName"`
	SecurityGroup         string  `json:"securityGroup"`
	IsSecurity            bool    `json:"isSecurity"`
	ShortCode             string  `json:"shortCode"`
	SignedChangePrice     float64 `json:"signedChangePrice"`
	SignedChangeRate      float64 `json:"signedChangeRate"`
	TotalMarketValue      float64 `json:"totalMarketValue"`
	TradePrice            float64 `json:"tradePrice"`
	TradeStrength         float64 `json:"tradeStrength"`
	TradeTime             string  `json:"tradeTime"`
}

const quotApi = "https://stockplus.com/api/securities.json"

func FetchSecuritiesByCodes(codes []string) (*FetchedSecuritiesResult, error) {
	resp, err := http.Get(quotApi + "?ids=" + strings.Join(codes, ","))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := FetchedSecuritiesResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

var TrendCodes = []string{
	"KOREA-D0011001", "KOREA-E4012001",
	"USA-DJI", "USA-COMP", "JAPAN-NI225", "SHANGHAI-000001", "GERMAN-DAX30", "BRITISH-FTSE100",
}

func FetchTrends(codes []string) (*FetchedSecuritiesResult, error) {
	var targetCodes []string
	if len(codes) > 0 {
		targetCodes = codes
	} else {
		targetCodes = TrendCodes
	}

	resp, err := http.Get(quotApi + "?ids=" + strings.Join(targetCodes, ","))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := FetchedSecuritiesResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}