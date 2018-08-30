package main

import (
	"net/http"
	"crawler/frontend/controller"
)

func main() {

	//不加这个会导致html中引用到的css不会handle
	http.Handle("/", http.FileServer(
		http.Dir("crawler/frontend/view/")))
/*	handler := controller.SearchResultHandler{}
	http.HandleFunc("/search", handler.ServeHTTP)*/

	http.Handle("/search", controller.CreateSearchResultHandler("crawler/frontend/view/template.html"))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
