package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {

	e := &engine.ConcurrentEngine{
		//Scheduler:  &scheduler.QueuedScheduler{},
		Scheduler:  &scheduler.SimpleScheduler{},
		WokerCount: 100,
	}

	e.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		})

}
