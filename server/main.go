package server

import (
	"net/http"

	"github.com/NovemberFoxtrot/punkt/templator"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templator.ThePool.Fill("index", "templates/layout.html", "templates/index.html")
	templator.ThePool.Pools["index"].Execute(w, nil)
}
