package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var thePort = flag.String("port", "8080", "port")

	flag.Parse()

	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	SetTemplates(Views)

	http.HandleFunc("/", Index)

	http.Handle("/touch", http.NotFoundHandler())

	fileServer := LoggingFunc(http.FileServer(http.Dir(wd + `/public`)))

	http.Handle(`/public/`, http.StripPrefix(`/public/`, fileServer))

	http.ListenAndServe(":"+*thePort, nil)
}
