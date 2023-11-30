package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ardanlabs/conf/v3"
	myLog "github.com/eduardobcolombo/colombostips/config-02/pkg/log"
	"github.com/eduardobcolombo/colombostips/config-02/pkg/newrelic"
)

func main() {
	exit, err := run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(exit)
}

func run() (int, error) {

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
			return 1, nil
		}
		return 1, fmt.Errorf("parsing config: %w", err)
	}

	logLevel := myLog.New(myLog.Config{
		Level: cfg.Log.Level,
	})
	fmt.Printf("\nUsing Log Level: %s\n\n", logLevel)

	_, err = newrelic.New(newrelic.Config{
		AppName:    cfg.NewRelic.AppName,
		LicenseKey: cfg.NewRelic.LicenseKey,
	})
	if err != nil {
		log.Fatal(err)
		return 1, err
	}

	return 0, nil
}
