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

	result := ParseCityList(body, "")

	const resultSize = 470

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
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
}
