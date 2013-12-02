package main

import (
	"flag"
	"fmt"
	"github.com/NovemberFoxtrot/punkt/fib"
	"github.com/NovemberFoxtrot/punkt/server"
	"net/http"
)

func main() {
	var theVersion = flag.Bool("version", false, "version")

	flag.Parse()

	if *theVersion == true {
		fmt.Println("0.0.1")
	}

	fmt.Println(fib.Fib(2))

	http.HandleFunc("/", server.Index)
	http.ListenAndServe(":8080", nil)
}
