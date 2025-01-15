package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var effects []string
	flag.Func("effects", "Rotation and Mirror", func(flagValue string) error {
		effects = strings.Split(flagValue, ",")
		return nil
	})
	flag.Parse()
	fmt.Println(effects)
}
