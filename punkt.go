package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/NovemberFoxtrot/punkt/server"
)

func main() {
	var theVersion = flag.Bool("version", false, "prints version")

	flag.Parse()

	if *theVersion == true {
		fmt.Println("0.0.1")
	}

	wd, _ := os.Getwd()

	http.HandleFunc("/", server.Index)

	http.Handle("/touch", http.NotFoundHandler())

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(wd+`/public`))))

	http.ListenAndServe(":8080", nil)
}
