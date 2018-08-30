package parser

import (
	"io/ioutil"
	"testing"
	"crawler/engine"
	"crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, e := ioutil.ReadFile("profile_test_data.html")
	if e != nil {
		panic(e)
	}

	result := parseProfile(contents, "http://album.zhenai.com/u/107194488", "霓裳")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/107194488",
		Type: "zhenai",
		Id:   "107194488",
		Payload: model.Profile{
			Name:         "霓裳",
			Age:          28,
			Height:       157,
			Marriage:     "未婚",
			Income:       "5001-8000元",
			Education:    "中专",
			Occupation:   "--",
			Gender:       "女",
			House:        "已购房",
			Car:          "已购车",
			WorkLocation: "上海徐汇区",
			Hukou:        "上海徐汇区",
			Xinzuo:       "--",
		},
	}

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}

}
