package main

import (
	"flag"
	"fmt"
	_ "github.com/NovemberFoxtrot/punkt/fib/"
	"github.com/NovemberFoxtrot/punkt/server/"
)

func main() {
	var theVersion = flag.Bool("version", false, "version")

	flag.Parse()

	if *theVersion == true {
		fmt.Println("0.0.1")
	}

	server.Init()
}
