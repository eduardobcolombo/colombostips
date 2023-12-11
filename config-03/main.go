package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/ardanlabs/conf/v3"
	"github.com/eduardobcolombo/colombostips/config-03/pkg/logger"
	"github.com/eduardobcolombo/colombostips/config-03/pkg/newrelic"
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
		Log struct {
			Level string `conf:"default:error"`
		}
		NewRelic struct {
			AppName    string `conf:"default:appName"`
			LicenseKey string `conf:"default:LicenceKey"`
		}
	}{}

	const prefix = ""
	// load env variables with conf lib.
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		// returning wrapped error for better context.
		return fmt.Errorf("parsing config: %w", err)
	}

	logCfg := logger.Config{
		Level: cfg.Log.Level,
	}

	logger := logger.New(logCfg)
	fmt.Println("Using Log Level:", logger.Level)

	nrCfg := newrelic.Config{
		AppName:    cfg.NewRelic.AppName,
		LicenseKey: cfg.NewRelic.LicenseKey,
	}

	// skipping the newrelic return for this demo
	if _, err = newrelic.New(nrCfg); err != nil {
		// returning wrapped error for better context.
		return fmt.Errorf("starting newrelic: %w", err)
	}

	return nil
}
