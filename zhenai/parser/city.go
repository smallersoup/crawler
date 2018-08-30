package parser

import (
	"crawler/engine"
		"regexp"
)

var proFileReg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var cityNextPageReg = regexp.MustCompile(`(http://www.zhenai.com/zhenghun/[^"]+)`)
	//<a href="http://album.zhenai.com/u/1361133512" target="_blank">怎么会迷上你</a>

func ParseCity(contents []byte, _ string) engine.ParserResult {

	submatch := proFileReg.FindAllSubmatch(contents, -1)

	//这里要把解析到的每个URL都生成一个新的request

	result := engine.ParserResult{}

	for _, m := range submatch {
		//name := string(m[2])
		//log.Printf("UserName:%s URL:%s\n", string(m[2]), string(m[1]))
		//数据库不需要城市信息和城市列表,只需要profile信息,所以这里不用append item
	/*	item := engine.Item{
			Payload:name,
		}*/

		//把用户信息人名加到item里
		//result.Items = append(result.Items, item)

		result.Requests = append(result.Requests,
			engine.Request{
				//用户信息对应的URL,用于之后的用户信息爬取
				Url: string(m[1]),
				//这个parser是对城市下面的用户的parse
				/*ParserFunc: func(bytes []byte) engine.ParserResult {
					//这里使用闭包的方式;这里不能用m[2],否则所有for循环里的用户都会共用一个名字
					//需要拷贝m[2] ---- name := string(m[2])
					return parseProfile(bytes, string(m[1]), name)
				},*/
				ParserFunc: ProfileParser(string(m[2])),
			})
	}

	matches := cityNextPageReg.FindAllSubmatch(contents, -1)


	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}

func ProfileParser(name string) engine.ParserFunc {
	return func(contents []byte, url string) engine.ParserResult {
		return parseProfile(contents, url, name)
	}
}
