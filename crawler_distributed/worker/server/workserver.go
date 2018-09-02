package main

import (
	"crawler/crawler_distributed/rpcsupport"
	"crawler/crawler_distributed/worker"
	"log"
	"fmt"
	"flag"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Println("must specify a port to listen on...")
		return
	}

	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), &worker.CrawlerService{}))
}
