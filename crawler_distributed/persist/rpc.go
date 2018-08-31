package persist

import (
	"crawler/engine"
	"github.com/olivere/elastic"
	"crawler/persist"
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
	}

	return err
}
