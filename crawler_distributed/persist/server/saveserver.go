package main

import (
	"crawler/crawler_distributed/persist"
	"crawler/crawler_distributed/rpcsupport"
	conf "crawler/crawler_distributed/config"
	"log"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/olivere/elastic.v5/config"
	"fmt"
	"flag"
)

var port = flag.Int("port", 0, "the port for me to listen on")

//ElasticSearch存储服务这样就写好了
func main() {
	flag.Parse()
	if *port == 0 {
		log.Println("must specify a port to listen on...")
		return
	}

	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), conf.Index))
}

func serveRpc(host string, index string) error {
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
		Client: client,
		Index:  index,
	})

}
