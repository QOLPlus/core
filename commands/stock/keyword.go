package stock

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SearchedAssetsResult struct {
	Keyword    string          `json:"keyword"`
	Asset      []SearchedAsset `json:"assets"`
	NextCursor string          `json:"nextCursor"`
}

type SearchedAsset struct {
	Type             int64  `json:"type"`
	Code             string `json:"type"`
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
