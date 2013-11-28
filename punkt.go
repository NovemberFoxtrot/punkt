package main

import (
	"flag"
	"fmt"
)

func main() {
	var theVersion = flag.Bool("version", false, "print version")

	flag.Parse()

	if *theVersion == true {
		fmt.Println("0.0.1")
	}
}
