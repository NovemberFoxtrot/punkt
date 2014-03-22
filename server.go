package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/iwanbk/gobeanstalk"
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
		ThePool.Fill(view.Index, view.Layout, view.Content)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	translations := map[string]string{}

	if strings.HasPrefix(r.Header.Get("Accept-Language"), "ja") == true {
		translations["greeting"] = "今日は"
	} else {
		translations["greeting"] = "Hello"
	}

	if r.Method == "POST" {
		fmt.Println(r.FormValue("email"))

		conn, err := gobeanstalk.Dial("localhost:11300")

		if err != nil {
			log.Fatal(err)
		}

		id, err := conn.Put([]byte(r.FormValue("Name")), 0, 0, 10)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Job id %d inserted\n", id)
	}

	ThePool.Pools["index"].Execute(w, translations)
}
