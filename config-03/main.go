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
		log.Fatal(err)
	}
}

func run() error {
	cfg := struct {
		Log struct {
			Level string `conf:"default:error"`
		}
		NewRelic struct {
			AppName    string `conf:"default:appName"`
			LicenseKey string `conf:"default:LicenceKey"`
		}
	}{}

	const prefix = ""
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
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
	fmt.Printf("\nUsing Log Level: %s\n\n", logger.Level)

	nrCfg := newrelic.Config{
		AppName:    cfg.NewRelic.AppName,
		LicenseKey: cfg.NewRelic.LicenseKey,
	}

	// skipping the newrelic return for this demo
	if _, err = newrelic.New(nrCfg); err != nil {
		return err
	}

	return nil
}
