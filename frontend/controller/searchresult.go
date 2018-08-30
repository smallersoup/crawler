package controller

import (
	"crawler/frontend/view"
	"github.com/olivere/elastic"
	"net/http"
	"strings"
	"strconv"
	"github.com/olivere/elastic/config"
	"crawler/frontend/model"
	"context"
	"reflect"
	"crawler/engine"
	"regexp"
	"log"
)

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

	from := currentPage * 10

	log.Printf("query string q=%s\n", q)
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

	log.Printf("共%d页;当前%d页\n", result.TotalPage, result.CurrentPage)

	return result, nil
}

func CreateSearchResultHandler(templateName string) SearchResultHandler {

	//sniff用来维护客户端和集群状态,但是集群运行在docker中,内网中无法sniff,以下这样默认是http://127.0.0.1:9200，即elasticsearch在本机上
	//client, err := elastic.NewClient(elastic.SetSniff(false))
	sniff := false
	client, err := elastic.NewClientFromConfig(
		&config.Config{
			//URL:   "http://192.168.1.102:9200",
			URL:   "http://127.0.0.1:9200",
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
