package client

import (
	"crawler/engine"
			"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/worker"
		"net/rpc"
)

func CreateCrawlerServiceProcessor(clients chan *rpc.Client) (engine.Processor) {

	return func(request engine.Request) (engine.ParserResult, error) {

		//先把engine.Request序列化为可网络上传输的worker.Request
		serialRequest := worker.SerializeRuquest(request)

		//然后调用CrawlerServiceRpc and args is worker.Request and worker.ParserResult
		var result worker.ParserResult

		//不同的客户端worker从clientPool里取client调用rpc进行fetch和parse
		c := <- clients
		err := c.Call(config.CrawlerServiceRpc, serialRequest, &result)

		if err != nil {
			return engine.ParserResult{}, err
		}

		//将网络上传输回来的result反序列化为engine.ParserResult
		return worker.DeserializeResult(result), nil
	}
}
