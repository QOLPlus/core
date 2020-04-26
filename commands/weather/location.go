package weather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const searchLocationApi = "https://ac.weather.naver.com/ac?q_enc=utf-8&r_format=json&r_enc=utf-8&r_lt=1&st=1&q="

type RegionEntry struct {
	Name string
	Code string
}
type FetchedLocationList struct {
	Query []string          `json:"query"`
	Items [][1][2][1]string `json:"items"`
}
func (f FetchedLocationList) GetFirstRegionCode() *RegionEntry {
	if len(f.Items) == 0 {
		return nil
	}

	firstItem := f.Items[0][0]

	return &RegionEntry{
		Name: firstItem[0][0],
		Code: firstItem[1][0],
	}
}

func FetchLocationByKeyword(keyword string) (*FetchedLocationList, error) {
	resp, err := http.Get(searchLocationApi + keyword)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := FetchedLocationList{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

