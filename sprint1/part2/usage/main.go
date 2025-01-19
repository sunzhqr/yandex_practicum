package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1"

var MyUsage = func() {
	fmt.Fprintf(flag.CommandLine.Output(), "Version: %s\n", version)
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = MyUsage
	flag.Parse()
}
