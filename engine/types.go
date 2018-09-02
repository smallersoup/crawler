package engine

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string
	Url     string
	Type    string
	Payload interface{}
}

type ParserFunc func(contents []byte, url string) ParserResult

type FuncParser struct {
	parser ParserFunc
	funcName string
}

func (f *FuncParser) Parser(contents []byte, url string) ParserResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (funcName string, args interface{}) {
	return f.funcName, nil
}

func NewFuncParser(p ParserFunc, funcName string) *FuncParser {
	return &FuncParser{
		parser : p,
		funcName: funcName,
	}
}

type Parser interface {
	Parser(contents []byte, url string) ParserResult
	Serialize() (funcName string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type NilParser struct {}

func (NilParser) Parser(_ []byte, _ string) ParserResult {
	return ParserResult{}
}

func (NilParser) Serialize() (_ string, _ interface{}) {
	return "NilParser", nil
}




