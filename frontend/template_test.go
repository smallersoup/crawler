package frontend

import (
	"testing"
	"html/template"
	"learngo/crawler/model"
	common "learngo/crawler/frontend/model"
	"os"
	"learngo/crawler/engine"
)

func TestTemplate(t *testing.T) {

	template := template.Must(template.ParseFiles("template.html"))

	page := common.SearchResult{}

	if err := template.Execute(os.Stdout, page); err != nil {
		panic(err)
	}
}

func TestTemplate2(t *testing.T) {

	template := template.Must(template.ParseFiles("template.html"))

	page := common.SearchResult{}

	page.Hits = 123
	page.Start = 0
	item := engine.Item {
		Url:  "http://album.zhenai.com/u/107194488",
		Type: "zhenai",
		Id:   "107194488",
		Payload: model.Profile{
			Name:       "霓裳",
			Age:        28,
			Height:     157,
			Marriage:   "未婚",
			Income:     "5001-8000元",
			Education:  "中专",
			Occupation: "程序媛",
			Gender:     "女",
			House:      "已购房",
			Car:        "已购车",
			Hukou:      "上海徐汇区",
			Xingzuo:    "水瓶座",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}


	file, _ := os.Create("template_test.html")

	if err := template.Execute(file, page); err != nil {
		panic(err)
	}
}
