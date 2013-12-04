package main

import (
	"flag"
	"fmt"
	"github.com/NovemberFoxtrot/punkt/server"
	"net/http"
)

func main() {
	var theVersion = flag.Bool("version", false, "version")

	flag.Parse()

	if *theVersion == true {
		fmt.Println("0.0.1")
	}

	http.HandleFunc("/", server.Index)

  http.Handle("/touch", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
