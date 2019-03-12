package main

import (
	"distributedCrawler/engine"
	"distributedCrawler/scheduler"
	"distributedCrawler/zhenai/parser"
)

func main() {
	//simpleEngine :=engine.SimpleEngine{}
	//simpleEngine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:50,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}
