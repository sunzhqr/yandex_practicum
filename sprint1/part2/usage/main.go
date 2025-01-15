package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s v%s:\n", os.Args[0], version)
		flag.PrintDefaults()
	}
	flag.Parse()
}
