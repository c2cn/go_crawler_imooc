package engine

import (
	"distributedCrawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, v := range parseResult.Items {
			log.Printf("Get Item %v", v)
		}
	}
}

func worker(r Request) (ParseResult, error) {

	log.Printf("Fecthing %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetcher error at %s:%v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
