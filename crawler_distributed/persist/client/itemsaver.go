package client

import (
	"log"
	"crawler/crawler_distributed/rpcsupport"
	"crawler/engine"
	"crawler/crawler_distributed/config"
	)

func ItemSaver(host string) (chan engine.Item, error) {

	//start ItemSaveServiceClient
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		for {
			item := <-out
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil || result != "ok" {
				log.Printf("Save item error; item: %v Error: %v\n", item, err)
				continue
			}
		}
	}()

	return out, nil
}
