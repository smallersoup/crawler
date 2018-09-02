package main

import (
	"crawler/crawler_distributed/rpcsupport"
	"crawler/crawler_distributed/worker"
	"time"
	"learngo/crawler_distributed/config"
	"testing"
)

//only test WorkerServer return worker.ParserResult
func TestWorkerServer(t *testing.T) {
	const host = ":9000"

	go rpcsupport.ServeRpc(host, &worker.CrawlerService{})

	time.Sleep(1 * time.Second)

	client, err := rpcsupport.NewClient(host)

	if err != nil {
		panic(err)
	}

	req := worker.Request {
		Url: "http://album.zhenai.com/u/108666172",
		Parser: worker.SerializedParser {
			Args: "清宁",
			FuncName: config.ParseProfile,
		},
	}

	var result worker.ParserResult
	if err := client.Call("CrawlerService.Process", req, &result); err != nil {
		t.Error(err)
	}

	t.Logf("parseResult:%v\n", result)
}
