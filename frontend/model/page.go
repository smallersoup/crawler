package model

import "learngo/crawler/engine"

type SearchResult struct {
	Hits int
	Start int
	Items []engine.Item
}
