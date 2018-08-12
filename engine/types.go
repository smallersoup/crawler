package engine

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
