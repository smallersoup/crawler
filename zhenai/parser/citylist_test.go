package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	body, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(body)

	const resultSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedCitys := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requests;but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if url != result.Requests[i].Url {
			t.Errorf("expectedUrl is %s "+
				"but is %s", url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+
			"items;but had %d", resultSize, len(result.Items))
	}

	for i, city := range expectedCitys {
		if city != result.Items[i].(string) {
			t.Errorf("expectedCity is %s "+
				"but is %s", city, result.Items[i].(string))
		}
	}
}
