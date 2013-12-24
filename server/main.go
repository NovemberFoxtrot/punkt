package server

import (
	"net/http"

	"github.com/NovemberFoxtrot/punkt/templator"
)

type View struct {
	Index   string
	Layout  string
	Content string
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
	templator.ThePool.Pools["index"].Execute(w, nil)
}
