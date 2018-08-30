package engine

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type ParserFunc func(contents []byte, url string) ParserResult

type Item struct {
	Id      string
	Url     string
	Type    string
	Payload interface{}
}

type Parser interface {
	Parser(contents []byte, url string) ParserResult
	//Serialize() (name string, args interface{})
}

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
