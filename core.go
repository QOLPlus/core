package main

import (
	"fmt"
	"os"

	"github.com/QOLPlus/core/commands/stock"
)

func main() {
	if len(os.Args) < 2 {
		panic("Pass the keyword!")
	}

	keyword := os.Args[1]
	fmt.Println("Keyword:", keyword)

	searchAssetsResult, err := stock.SearchAssetsByKeyword(keyword)
	if err != nil {
		panic(err)
	}
	fmt.Println("Searched:", searchAssetsResult)

	codes := searchAssetsResult.FindExactlySameCodesByKeyword(keyword)
	if len(codes) == 0 {
		codes = searchAssetsResult.GetCodes()
		if len(codes) == 0 {
			panic("No codes!")
		}
	}

	fmt.Println("Codes:", codes)
	fetchSecuritiesResult, err := stock.FetchSecuritiesByCodes(codes)
	if err != nil {
		panic(err)
	}
	fmt.Println("Securities:", fetchSecuritiesResult)
}
