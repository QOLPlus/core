package stock

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SearchedAssetsResult struct {
	Keyword    string          `json:"keyword"`
	Assets     []SearchedAsset `json:"assets"`
	NextCursor string          `json:"nextCursor"`
}

type SearchedAsset struct {
	Type             string `json:"type"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	AssetId          string `json:"assetId"`
	DisplayedSubtype string `json:"displayedSubtype"`
	DisplayedCode    string `json:"displayedCode"`
}

func SearchAssetsByKeyword(keyword string) (*SearchedAssetsResult, error) {
	resp, err := http.Get("https://stockplus.com/api/search/assets?keyword=" + keyword)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := SearchedAssetsResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *SearchedAssetsResult) GetCodes() []string {
	var codes []string
	for _, asset := range r.Assets {
		codes = append(codes, asset.AssetId)
	}
	return codes
}

func (r *SearchedAssetsResult) FindExactlySameCodesByKeyword(keyword string) []string {
	var codes []string
	for _, asset := range r.Assets {
		if asset.Name == keyword {
			codes = append(codes, asset.AssetId)
			break
		}
	}
	return codes
}