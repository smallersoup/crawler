package model

type SearchResult struct {
	Hits int64
	Start int
	Query string
	PrevFrom int
	NextFrom int
	CurrentPage int
	TotalPage int64
	Items []interface{}
	//Items []engine.Item
}
