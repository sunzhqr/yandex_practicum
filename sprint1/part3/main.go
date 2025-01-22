package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Getenv(key string) string
	// get environment variable value by key
	user := os.Getenv("USER")
	fmt.Println(user)

	// os.Environ() []string for get all environment variable
	for idVar, envVar := range os.Environ() {
		fmt.Printf("%d: %s\n", idVar, envVar)
	}

	// os.LookupEnv(key string) (string, bool) for checking null value
	nothing, ok := os.LookupEnv("NOTHING")
	if !ok {
		fmt.Println("We don't have given environment variable")
		os.Exit(1)
	} else {
		fmt.Println(nothing)
	}

}
