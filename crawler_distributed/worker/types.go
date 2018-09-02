package worker

import (
	"crawler/engine"
	"crawler/crawler_distributed/config"
	"crawler/zhenai/parser"
	"fmt"
	"log"
)

//序列化后的Parser,方法名+参数,因为函数没法序列化后在网络上传输
//{"funcName" : "ParseCityList", "args", nil},{"funcName" : "ProfileParser" : "args" : "username"}
type SerializedParser struct {
	FuncName string
	Args     interface{}
}

//可序列化在网络上传输的Request
type Request struct {
	Url    string
	Parser SerializedParser
}

//可序列化在网络上传输的ParserResult
type ParserResult struct {
	Items    []engine.Item
	Requests []Request
}

//不可序列化的engine.Request转可序列化的Request
func SerializeRuquest(r engine.Request) Request {

	funcName, args := r.Parser.Serialize()

	req := Request{
		r.Url,
		SerializedParser{
			funcName,
			args,
		},
	}

	return req
}

func SerializeParserResult(result engine.ParserResult) ParserResult {

	//Items可序列化,无需转换
	parseResult := ParserResult{
		Items: result.Items,
	}

	//Requests需要range一个一个转换
	for _, req := range result.Requests {
		serialReq := SerializeRuquest(req)
		parseResult.Requests = append(parseResult.Requests, serialReq)
	}

	return parseResult
}

func DeserializeRequest(req Request) (engine.Request, error) {

	deserializeParser, err := DeserializeParser(req.Parser)

	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:    req.Url,
		Parser: deserializeParser,
	}, nil
}

//convert to ParserFunc by parser funcName
func DeserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.FuncName {
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg:%+v", p.Args)
		}
	default:
		return nil, fmt.Errorf("unknown parser func name %s", p.FuncName)
	}
}

func DeserializeResult(result ParserResult) engine.ParserResult {

	desParserResult := engine.ParserResult{
		Items: result.Items,
	}

	//Requests需要range一个一个转换
	for _, req := range result.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("DeserializeRequest %+v Error:%+v\n", request, err)
			continue
		}
		desParserResult.Requests = append(desParserResult.Requests, request)
	}

	return desParserResult
}
