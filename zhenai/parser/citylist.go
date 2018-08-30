package parser

import (
	"crawler/engine"
	"regexp"
)

var	cityListReg = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte, _ string) engine.ParserResult {

	submatch := cityListReg.FindAllSubmatch(contents, -1)

	//这里要把解析到的每个URL都生成一个新的request
	result := engine.ParserResult{}

	//limit := 10
	for _, m := range submatch {

		//数据库不需要城市信息和城市列表,只需要profile信息,所以这里不用append item
	/*	item := engine.Item{
			Payload:string(m[2]),
		}*/
		//把城市名字append到item里
		//result.Items = append(result.Items, item)

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
