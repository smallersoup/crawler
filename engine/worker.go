package engine

import (
	"log"
	"crawler/fetcher"
)

func worker(r Request) (ParserResult, error) {

	//log.Printf("fetching url:%s\n", r.Url)
	//爬取数据
	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("fetch url: %s; err: %v\n", r.Url, err)
		//发生错误继续爬取下一个url
		return ParserResult{}, err
	}

	//解析爬取到的结果
	return r.ParserFunc(body, r.Url), nil
}
