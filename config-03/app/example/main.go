package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/ardanlabs/conf/v3"
	"github.com/eduardobcolombo/colombostips/config-03/foundation/logger"
	"github.com/eduardobcolombo/colombostips/config-03/foundation/newrelic"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var cfg struct {
		Log struct {
			Level string `conf:"default:error"`
		}
		NewRelic struct {
			AppName    string `conf:"default:appName"`
			LicenseKey string `conf:"default:LicenceKey"`
		}
	}

	const prefix = ""
	if help, err := conf.Parse(prefix, &cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}

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

	if _, err := newrelic.New(nrCfg); err != nil {
		return fmt.Errorf("starting newrelic: %w", err)
	}

	return nil
}
