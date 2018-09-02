package persist

import (
	"crawler/engine"
	"crawler/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

//存储服务
type ItemSaveService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaveService) Save (item engine.Item, result *string) error {

	_, err := persist.Save(s.Client, s.Index, item)

	if err == nil {
		*result = "ok"
		log.Printf("Item: %v ;Saved", item)
	} else {
		log.Printf("Save item %v Error: %v", item, err)
	}

	return err
}
