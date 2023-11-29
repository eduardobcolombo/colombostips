package main

import (
	"fmt"
	"log"
	"os"

	myLog "github.com/eduardobcolombo/colombostips/config-02/pkg/log"
	"github.com/eduardobcolombo/colombostips/config-02/pkg/newrelic"
	"github.com/kelseyhightower/envconfig"
)

// Config represents a configuration type to receive env variables
type Config struct {
	Log      myLog.Config
	NewRelic newrelic.Config `envconfig:"NEW_RELIC" desc:"NewRelic config"`
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
		_ = envconfig.Usage("", &cfg)
		log.Fatalf("error processing envconfig %v", err)
		return 1, err
	}

	logLevel := myLog.New(cfg.Log)
	fmt.Printf("\nUsing Log Level: %s\n\n", logLevel)

	_, err := newrelic.New(cfg.NewRelic)
	if err != nil {
		log.Fatal(err)
		return 1, err
	}

	return 0, nil
}
