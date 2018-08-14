package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
		"log"
	"net/http"
	_ "net/http/pprof"
	"crawler/scheduler"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	e := &engine.ConcurrentEngine{
		//Scheduler: &scheduler.QueuedScheduler{},

		Scheduler:  &scheduler.SimpleScheduler{},
		WokerCount: 100,
	}

	/*	e.Run(
			engine.Request{
				Url:        "http://www.zhenai.com/zhenghun",
				ParserFunc: parser.ParseCityList,
			})*/

	e.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun/shanghai",
			ParserFunc: parser.ParseCityList,
		})

}
