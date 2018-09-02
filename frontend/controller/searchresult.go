package controller

import (
	"crawler/frontend/view"
	"net/http"
	"strings"
	"strconv"
	"crawler/frontend/model"
	"context"
	"reflect"
	"crawler/engine"
	"regexp"
	"log"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/olivere/elastic.v5/config"
	)

const size = 10
const es_max_result_window = 10000

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

//localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	q := strings.TrimSpace(req.FormValue("q"))

	currentPage, err := strconv.Atoi(req.FormValue("current"))

	if err != nil {
		currentPage = 1
	}

	//fmt.Fprintf(w, "q=%s, from=%d\n", q, from)

	page, err := h.getSearchResult(q, currentPage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h SearchResultHandler) getSearchResult(q string, currentPage int) (model.SearchResult, error) {

	var result model.SearchResult
	result.Query = q

	//然后重写查询条件
	q = rewriteQueryString(q)

	from := (currentPage - 1) * 10

	//ElasticSearch默认max_result_window为10000,超出范围会报错,也可以通过以下curl修改window大小,但是会增加内存和cpu开销,项目中需要权衡
	// curl -XPUT http://127.0.0.1:9200/dating_profile/_settings -d '{ "index" : { "max_result_window" : 100000000}}'
	//这里简单做校验
	/*if from + size > es_max_result_window {
		return result, fmt.Errorf("ElasticSearch setting required:Result window is too large, from + size must be less than or equal to: [10000], but %d\n", from + size)
	}*/

	log.Printf("query string q=%s, from=%d\n", q, from)
	resp, err := h.client.Search("dating_profile").Query(elastic.NewQueryStringQuery(q)).
		From(from).Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	//result.PrevFrom = result.Start - len(result.Items)
	//result.NextFrom = result.Start + len(result.Items)
	result.CurrentPage = currentPage

	if result.Hits % 10 > 0 {
		result.TotalPage = result.Hits / 10 + 1
	} else {
		result.TotalPage = result.Hits / 10
	}

	log.Println("-----------------------------------------")
	log.Printf("共%d页;当前%d页\n", result.TotalPage, result.CurrentPage)
	//log.Printf("共%d页;当前%d页\n", result.TotalPage, result.CurrentPage)

	return result, nil
}

func CreateSearchResultHandler(templateName string) SearchResultHandler {

	//sniff用来维护客户端和集群状态,但是集群运行在docker中,内网中无法sniff,以下这样默认是http://127.0.0.1:9200，即elasticsearch在本机上
	//client, err := elastic.NewClient(elastic.SetSniff(false))
	sniff := false
	client, err := elastic.NewClientFromConfig(
		&config.Config{
			//URL:   "http://192.168.1.102:9200",
			URL:   "http://192.168.1.101:9200",
			Sniff: &sniff,
		})

	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(templateName),
		client: client,
	}
}

func rewriteQueryString(q string) string {

	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	//$1代表上面的组
	return re.ReplaceAllString(q, "Payload.$1:")
}
