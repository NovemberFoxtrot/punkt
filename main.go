package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

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
	t := render("templates/layout.html", "templates/about.html")
	t.Execute(w, nil)
}

type entry struct {
	Title string
	Body  string
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := []entry{
		{"title", "body"},
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

	t.Execute(w, data)
}

func Privacy(w http.ResponseWriter, r *http.Request) {
	t := render("templates/layout.html", "templates/privacy.html")
	t.Execute(w, nil)
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
	http.HandleFunc("/privacy", Privacy)
	http.HandleFunc("/", Index)

	http.Handle("/touch", http.NotFoundHandler())

	fileServer := LoggingFunc(http.FileServer(http.Dir(wd + `/public`)))

	http.Handle(`/public/`, http.StripPrefix(`/public/`, fileServer))

	http.ListenAndServe(":"+*thePort, nil)
}
