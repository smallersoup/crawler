package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
	//_ "net/http/pprof"
	"crawler/scheduler"
	"crawler/persist"
	"os"
	"fmt"
	"log"
	"crawler/crawler_distributed/config"
)

func init() {
	outfile, err := os.OpenFile("root.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开
	if err != nil {
		fmt.Println(*outfile, "open failed")
		os.Exit(1)
	}

	log.SetOutput(outfile)  //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("---------------Start Crawler----------------") //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了
}

func main() {

	/*	go func() {
			log.Println(http.ListenAndServe("localhost:8080", nil))
		}()*/

	itemChan, err := persist.ItemSaver("dating_profile")

	if err != nil {
		panic(err)
	}

	e := &engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},

		//Scheduler:  &scheduler.SimpleScheduler{},
		WokerCount: 100,
		ItemChan:   itemChan,
		RequestProcessor: engine.Worker,
	}

	/*	e.Run(
			engine.Request{
				Url:        "http://www.zhenai.com/zhenghun",
				ParserFunc: parser.ParseCityList,
			})*/

	e.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
		})
}
