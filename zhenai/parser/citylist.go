package parser

import (
	"crawler/engine"
	"regexp"
)

var	cityListReg = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParserResult {

	submatch := cityListReg.FindAllSubmatch(contents, -1)

	//这里要把解析到的每个URL都生成一个新的request

	result := engine.ParserResult{}

	//limit := 10
	for _, m := range submatch {
		//把城市名字append到item里
		result.Items = append(result.Items, string(m[2]))

		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				//这个parser是对城市下面的用户的parse
				ParserFunc: ParseCity,
			})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}

	return result
}
