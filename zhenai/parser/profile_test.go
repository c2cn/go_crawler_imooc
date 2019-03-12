package parser

import (
	"distributedCrawler/fetcher"
	"fmt"
	"testing"
)

func TestProfile(t *testing.T) {
	contents, _ := fetcher.Fetch("http://album.zhenai.com/u/1441118068")

	fmt.Printf("%s\n", contents)
	fmt.Println(ParseProfile(contents, "haha"))

}
