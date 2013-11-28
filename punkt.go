package main

import (
	"flag"
	"fmt"
	_ "fib"
	"server"
)

func main() {
	var theVersion = flag.Bool("version", false, "print version")

	flag.Parse()

	if *theVersion == true {
		fmt.Println("0.0.1")
	}

	server.Init()
}
