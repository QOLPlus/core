package weather

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

const weatherUrl = "https://m.weather.naver.com/m/main.nhn?regionCode="

type FetchWeatherResult struct {
	Location 		   string
	Temperature        float64
	TemperatureDayLow  float64
	TemperatureDayHigh float64
	TemperatureDayFeel float64
	DiffWithYesterday  float64
	Status             string
}

func FetchWeather(region *RegionEntry) (*FetchWeatherResult, error) {
	resp, err := http.Get(weatherUrl + region.Code)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	result := FetchWeatherResult{
		Location: doc.Find(".section_location a.title strong").Text(),
	}
	err = parseTemperature(&result, doc)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func parseTemperature(result *FetchWeatherResult, doc *goquery.Document) error {
	result.Temperature, _ = strconv.ParseFloat(doc.Find(".section_content .current").Text(), 64)
	result.TemperatureDayLow, _ =  strconv.ParseFloat(doc.Find(".section_content .day .day_low .degree_code").Text(), 64)
	result.TemperatureDayHigh, _ = strconv.ParseFloat(doc.Find(".section_content .day .day_high .degree_code").Text(), 64)
	result.TemperatureDayFeel, _ = strconv.ParseFloat(doc.Find(".section_content .day .day_feel .degree_code").Text(), 64)

	summaryDom := doc.Find(".section_content .weather_set_summary")
	summaryHtml, err := summaryDom.Html()
	if err != nil {
		return err
	}

	diffDirection := 1.0
	if strings.Contains(summaryHtml, "낮아") {
		diffDirection = -1.0
	}
	diffWithYesterday, _ := strconv.ParseFloat(summaryDom.Find(".degree.degree_code").Text(), 64)
	result.DiffWithYesterday = diffWithYesterday * diffDirection
	result.Status = strings.SplitN(summaryHtml, "<br/>", 2)[0]

	return nil
}
