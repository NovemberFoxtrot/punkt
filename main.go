package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/NovemberFoxtrot/punkt/server"
)

func main() {
	var thePort = flag.String("port", "8080", "port")

	flag.Parse()

	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	server.SetTemplates(server.Views)

	http.HandleFunc("/", server.Index)

	http.Handle("/touch", http.NotFoundHandler())

	fileServer := server.LoggingFunc(http.FileServer(http.Dir(wd + `/public`)))

	http.Handle(`/public/`, http.StripPrefix(`/public/`, fileServer))

	http.ListenAndServe(":"+*thePort, nil)
}
