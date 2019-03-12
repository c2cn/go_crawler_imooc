package parser

import (
	"distributedCrawler/fetcher"
	"fmt"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	fmt.Printf("%s\n", contents)

	ParseCityList(contents)
}
