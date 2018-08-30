package persist

import (
	"crawler/model"
	"testing"
		"gopkg.in/olivere/elastic.v5"
	"gopkg.in/olivere/elastic.v5/config"
	"context"
	"encoding/json"
	"crawler/engine"
)

func TestSave(t *testing.T) {

	expectedItem := engine.Item{
		Id: "108666172",
		Type: "zhenai",
		Url: "http://album.zhenai.com/u/108666172",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hukou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	sniff := false
	//TODO: Try to startup elastic search here using docker go client
	client, err := elastic.NewClientFromConfig(
		&config.Config{
			URL:   "http://192.168.1.102:9200",
			Sniff: &sniff,
		})
	const index = "dating_test"
	id, err := save(client, index, expectedItem)

	if err != nil {
		panic(err)
	}

	t.Logf("Id:%s\n", id)

	itemService := client.Get().
		Index("dating_profile").
		Type(expectedItem.Type).
		Id(expectedItem.Id)

	resp, err := itemService.Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%v\n", resp)
	t.Logf("%+v\n", resp)

	t.Logf("%s\n", *resp.Source)

	actualProfile := engine.Item{}

	err = json.Unmarshal(*resp.Source, &actualProfile)

	if err != nil {
		panic(err)
	}

	//地址是不一样的
	t.Logf("actualProfile=%p;expectedProfile%p\n", &actualProfile, &expectedItem)

	//actualProfile中的payload是map类型,而expectedItem的payload是profile类型,需要转换
	actualProfile.Payload, _ = model.FromJsonObj(actualProfile.Payload)

	//但是struct对象里内容是一样的
	if actualProfile != expectedItem {
		 t.Errorf("Got %v, expected %v\n", actualProfile, expectedItem)
	}
}
