package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	if err := run(); err != nil {
		// Fatal will always call os.Exit(1).
		log.Fatal(err)
	}
}

func run() error {
	// cfg variable with zero value construction.
	var cfg = struct {
		LogLevel string `split_words:"true" default:"debug"`
	}{}

	// load env variables with envconfig.
	if err := envconfig.Process("", &cfg); err != nil {
		// returning wrapped error for better context.
		return fmt.Errorf("loading env config: %w", err)
	}

	// Print out the log level from env variable.
	fmt.Println("LogLevel: ", cfg.LogLevel)

	return nil
}
