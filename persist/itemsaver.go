package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"gopkg.in/olivere/elastic.v5/config"
	. "crawler/engine"
	"errors"
)

func ItemSaver(index string) (chan Item, error) {

	//sniff用来维护客户端和集群状态,但是集群运行在docker中,内网中无法sniff,以下这样默认是http://127.0.0.1:9200，即elasticsearch在本机上
	//client, err := elastic.NewClient(elastic.SetSniff(false))
	sniff := false
	client, err := elastic.NewClientFromConfig(
		&config.Config{
			URL:   "http://192.168.1.101:9200",
			Sniff: &sniff,
		})

	if err != nil {
		return nil, err
	}

	out := make(chan Item)

	go func() {

		itemCount := 0

		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			_, err := Save(client, index, item)
			if err != nil {
				log.Printf("Save item error; item: %v Error: %v\n", item, err)
				continue
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, index string, item Item) (id string, err error) {

	if item.Type == "" {
		return "", errors.New("must supply Type")
	}

	//Index=dating_profile相当于数据名;Type=zhenai相当于表名;ID可以指定,也可以不指定
	resp, err := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	//log.Println(resp)
	return resp.Id, nil
}
