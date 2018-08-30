package view

import (
	"html/template"
	"io"
	"crawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

//减法，为了在模板里用减1
func Sub(a, b int) int {
	return a - b
}

//加法，为了在模板里用加1
func Add(a, b int) int {
	return a + b
}

func CreateSearchResultView(filename string) SearchResultView {

	funcMaps := template.FuncMap{"Add": Add, "Sub": Sub}
	//template := template.Must(template.ParseFiles(filename)).Funcs(funcMaps)
	template, err := template.New("template.html").Funcs(funcMaps).ParseFiles(filename)

	if err != nil {
		panic(err)
	}

	return SearchResultView{
		template,
	}
}

func (s SearchResultView) Render (w io.Writer, data model.SearchResult) error {

	return s.template.Execute(w, data)
}
