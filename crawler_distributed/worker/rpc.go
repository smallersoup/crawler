package worker

import (
	"crawler/engine"
	)

type CrawlerService struct {
}

func (c *CrawlerService) Process(req Request, result *ParserResult) error {

	//先将序列化后的Request转换
	request, err := DeserializeRequest(req)

	if err != nil {
		return err
	}

	parseResult, err := engine.Worker(request)

	if err != nil {
		return err
	}

	//最后再转换为可序列化的ParserResult
	*result = SerializeParserResult(parseResult)
	return nil
}
