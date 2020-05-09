package weather

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

type FetchWeatherResult struct {
	Location string
	Status   string

	// 온도
	Temperature        float64
	TemperatureDayLow  float64
	TemperatureDayHigh float64
	TemperatureDayFeel float64

	// 미세먼지
	FineDust            int
	FineDustStatus      string
	UltraFineDust       int
	UltraFineDustStatus string

	// 자외선
	UltravioletLay       int
	UltravioletLayStatus string
}

const weatherUrl = "https://n.weather.naver.com/today/"

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
		Location: doc.Find(".location_area .location_name").Text(),
	}

	err = parseTemperature(&result, doc)
	if err != nil {
		return nil, err
	}

	err = parseDust(&result, doc)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func parseTemperature(result *FetchWeatherResult, doc *goquery.Document) error {
	current := doc.Find(".today_weather .weather_area .current")
	result.Temperature = parseOnlyTemperature(current)

	todayHigh := doc.Find(".today_weather .weather_area .degree_group .degree_height")
	result.TemperatureDayHigh = parseOnlyTemperature(todayHigh)

	todayLow := doc.Find(".today_weather .weather_area .degree_group .degree_low")
	result.TemperatureDayLow = parseOnlyTemperature(todayLow)

	todayFeel := doc.Find(".today_weather .weather_area .degree_group .degree_feel")
	result.TemperatureDayFeel = parseOnlyTemperature(todayFeel)

	summary := doc.Find(".today_weather .weather_area .summary")
	summaryHtml, err := summary.Html()
	if err != nil {
		return err
	}
	if strings.Contains(summaryHtml, "<br/>") {
		summarySplit := strings.SplitN(summaryHtml, "<br/>", 2)
		result.Status = summarySplit[1]
	} else {
		result.Status = summary.Text()
	}

	return nil
}

func parseOnlyTemperature(s *goquery.Selection) float64 {
	parsed := ""

	s.Contents().Each(func(i int, selection *goquery.Selection) {
		if goquery.NodeName(selection) == "#text" {
			parsed = strings.ReplaceAll(selection.Text(), "°", "")
			return
		}
	})

	temperature, _ := strconv.ParseFloat(parsed, 64)
	return temperature
}

func parseDust(result *FetchWeatherResult, doc *goquery.Document) error {
	items := doc.Find(".today_weather .today_chart_list .item_today")
	items.Each(func(i int, selection *goquery.Selection) {
		switch selection.Find("strong.ttl").Text() {
		case "미세먼지":
			result.FineDustStatus = selection.Find(".level_text").Text()
			result.FineDust, _ = strconv.Atoi(selection.Find(".chart .value").Text())
		case "초미세먼지":
			result.UltraFineDustStatus = selection.Find(".level_text").Text()
			result.UltraFineDust, _ = strconv.Atoi(selection.Find(".chart .value").Text())
		case "자외선":
			result.UltravioletLayStatus = selection.Find(".level_text").Text()
			result.UltravioletLay, _ = strconv.Atoi(selection.Find(".chart .value").Text())
		default:
		}
	})

	return nil
}