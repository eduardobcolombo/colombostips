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
	cfg := struct {
		LogLevel string `split_words:"true" default:"debug"`
	}{}

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("error processing envconfig %v", err)
		return err
	}

	fmt.Printf("LogLevel: %s\n", cfg.LogLevel)

	return nil
}
