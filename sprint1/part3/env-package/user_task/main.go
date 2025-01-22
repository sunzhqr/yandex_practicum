package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	User string `env:"USER"`
}

func (c Config) String() string {
	return string(c.User)
}

func main() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
}
