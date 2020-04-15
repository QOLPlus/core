package main

import "fmt"
import "github.com/QOLPlus/core/commands/stock"

func main() {
	fmt.Println("Run!")

	searchAssetsResult, err := stock.SearchAssetsByKeyword("네이")
	if err != nil {
		panic(err)
	}
	fmt.Println(searchAssetsResult)

	fmt.Println("==================")

	codes := []string{"KOREA-A035720", "KOREA-D0011001"}
	fetchSecuritiesResult, err := stock.FetchSecuritiesByCodes(codes)
	if err != nil {
		panic(err)
	}
	fmt.Println(fetchSecuritiesResult)

	fmt.Println("End!")
}
