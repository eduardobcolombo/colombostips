package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var cfg struct {
		LogLevel string `split_words:"true" default:"debug"`
	}

	const prefix = ""
	if err := envconfig.Process(prefix, &cfg); err != nil {
		return fmt.Errorf("loading env config: %w", err)
	}

	fmt.Println("LogLevel: ", cfg.LogLevel)

	return nil
}
