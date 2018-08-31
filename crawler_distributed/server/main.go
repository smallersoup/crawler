package main

import (
	"github.com/olivere/elastic"
	"github.com/olivere/elastic/config"
	"crawler/crawler_distributed/persist"
	"crawler/crawler_distributed/rpcsupport"
	conf "crawler/crawler_distributed/config"
	"log"
)

//ElasticSearch存储服务这样就写好了
func main() {

	log.Fatal(serveRpc(conf.Host, conf.Index))

}

func serveRpc(host, index string) error {
	//sniff用来维护客户端和集群状态,但是集群运行在docker中,内网中无法sniff,以下这样默认是http://127.0.0.1:9200，即elasticsearch在本机上
	//client, err := elastic.NewClient(elastic.SetSniff(false))
	sniff := false
	client, err := elastic.NewClientFromConfig(
		&config.Config{
			URL:   conf.EsUrl,
			Sniff: &sniff,
		})

	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Client:client,
		Index: index,
	})

}



