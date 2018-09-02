package main

import (
	itemSaverClient "crawler/crawler_distributed/persist/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"crawler/crawler_distributed/config"
			"log"
	workerClient "crawler/crawler_distributed/worker/client"
	"flag"
	"net/rpc"
	"crawler/crawler_distributed/rpcsupport"
	"strings"
)

/*func init() {
	outfile, err := os.OpenFile("root.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开
	if err != nil {
		fmt.Println(*outfile, "open failed")
		os.Exit(1)
	}

	log.SetOutput(outfile)  //设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("Start crawler..............") //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了
}*/

var (
	//multi worker
	workerPort = flag.String("workerport", "", "worker port (comma separated)")

	//single saver
	saverPort = flag.String("saverport", "", "saver port")
)

func main() {

	flag.Parse()
	/*	go func() {
			log.Println(http.ListenAndServe("localhost:8080", nil))
		}()*/

	//connect saveRpc
	itemChan, err := itemSaverClient.ItemSaver(*saverPort)

	if err != nil {
		panic(err)
	}

	//这里先不做参数正确性验证
	pool := createClientPool(strings.Split(*workerPort, ","))
	//connect workRpc
	processor := workerClient.CreateCrawlerServiceProcessor(pool)

	e := &engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},

		//Scheduler:  &scheduler.SimpleScheduler{},
		WokerCount: 100,
		ItemChan:   itemChan,
		RequestProcessor : processor,
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

func createClientPool(hosts []string) chan *rpc.Client {

	var clients []*rpc.Client

	//连接worker的client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)

		if err == nil {
			clients = append(clients, client)
			log.Printf("Connect port %s success!\n", h)
		} else {
			log.Printf("Failed connect port %s Error:%v\n", h, err)
		}
	}

	clientChan := make(chan *rpc.Client)

	//不断的将clientChan送进chan,让不100个客户端worker去抢
	go func() {
		for {
			for _, c := range clients {
				clientChan <- c
			}
		}
	}()

	return clientChan
}
