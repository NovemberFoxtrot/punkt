package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/iwanbk/gobeanstalk"
)

func LoggingFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(*r.URL)
		h.ServeHTTP(w, r)
	})
}

func render(filenames ...string) *template.Template {
	t := template.New("layout")

	t, err := t.ParseFiles(filenames...)

	if err != nil {
		log.Fatal("parsing:", err)
	}

	return t
}

func About(w http.ResponseWriter, r *http.Request) {
	render("templates/layout.html", "templates/about.html").Execute(w, nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

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

		fmt.Printf("Job %d inserted\n", id)
	}

	t := render("templates/layout.html", "templates/index.html")

	t.Execute(w, translations)
}

func main() {
	var thePort = flag.String("port", "8080", "port")

	flag.Parse()

	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	http.HandleFunc("/about", About)
	http.HandleFunc("/", Index)

	http.Handle("/touch", http.NotFoundHandler())

	// fileServer := LoggingFunc(http.FileServer(http.Dir(wd + `/public`)))

	fileServer := http.FileServer(http.Dir(wd + `/public`))

	http.Handle(`/public/`, http.StripPrefix(`/public/`, fileServer))

	http.ListenAndServe(":"+*thePort, nil)
}
