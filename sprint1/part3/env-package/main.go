package main

import (
	"log"
	"time"

	"github.com/caarlos0/env"
)

type Config struct {
	Files []string `env:"FILES" envSeparator:":"`
	Home  string   `env:"HOME"`
	// required требует, чтобы переменная TASK_DURATION была определена
	TaskDuration time.Duration `env:"TASK_DURATION,required"`
}

func main() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)
}
