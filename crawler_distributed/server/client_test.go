package main

import (
	"testing"
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"log"
	"time"
)

func TestClient(t *testing.T) {

	//start ItemSaveServiceServer

	go serveRpc(config.Host, config.Index)

	//等待一秒等ItemSaveServiceServer服务起来后再call
	time.Sleep(1 * time.Second)

	//start ItemSaveServiceClient
	client, err := rpcsupport.NewClient(config.Host)

	if err != nil {
		panic(err)
	}

	//Call save
	item := engine.Item{
		//Id:   "108666172",
		Type: "zhenai",
		Url:  "http://album.zhenai.com/u/108666172",
		Payload: model.Profile{
			Age:        20,
			Height:     115,
			Weight:     68,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "sap",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hukou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	result := ""
	err = client.Call("ItemSaveService.Save", item, &result)

	if err != nil || result != "ok"{
		t.Fatal(err)
	}

	log.Printf("result:%s\n", result)
}
