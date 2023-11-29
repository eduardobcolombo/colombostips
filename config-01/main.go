package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Config represents a configuration type to receive env variables
type Config struct {
	LogLevel string `split_words:"true" default:"debug"`
}

func main() {
	exit, err := run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(exit)
}

func run() (int, error) {

	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("error processing envconfig %v", err)
		return 1, err
	}

	fmt.Printf("LogLevel: %s\n", cfg.LogLevel)

	return 0, nil
}
