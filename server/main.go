package server

import (
	"fmt"
	"net/http"

	"github.com/NovemberFoxtrot/punkt/templator"
)

type View struct {
	Index   string
	Layout  string
	Content string
}

func LoggingFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(*r.URL)
		h.ServeHTTP(w, r)
	})
}

var Views = []View{
	{"index", "templates/layout.html", "templates/index.html"},
}

func SetTemplates(views []View) {
	for _, view := range views {
		templator.ThePool.Fill(view.Index, view.Layout, view.Content)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.URL)

	if r.Method == "POST" {
		fmt.Println(r.FormValue("Name"))
	}

	data := map[string]interface{}{"dude": []interface{}{1}}

	templator.ThePool.Pools["index"].Execute(w, data)
}
