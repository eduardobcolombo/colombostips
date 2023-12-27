package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ardanlabs/conf/v3"
	"github.com/eduardobcolombo/colombostips/config-04/logger"
	mynewrelic "github.com/eduardobcolombo/colombostips/config-04/newrelic"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type App struct {
	*logger.Logger
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var cfg = struct {
		Log struct {
			Level string `conf:"default:error"`
		}
		NewRelic struct {
			AppName    string `conf:"default:appName"`
			LicenseKey string `conf:"default:LicenceKey"`
		}
		Port string `conf:"default:8000"`
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
	defer logger.Log.Sync()

	nrCfg := mynewrelic.Config{
		AppName:    cfg.NewRelic.AppName,
		LicenseKey: cfg.NewRelic.LicenseKey,
	}

	nr, err := mynewrelic.New(nrCfg)
	if err != nil {
		return fmt.Errorf("initiating newrelic: %w", err)
	}

	app := App{Logger: logger}

	txn := nr.App.StartTransaction("Example transaction")
	defer txn.End()

	txn.AddAttribute("traceID", txn.GetTraceMetadata().TraceID)

	http.HandleFunc(newrelic.WrapHandleFunc(nr.App, "/", app.hello))

	port := fmt.Sprintf(":%s", cfg.Port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (a App) hello(w http.ResponseWriter, req *http.Request) {
	// Write in the response
	fmt.Fprintf(w, "hello\n")

	// Write in the log
	a.Log.Debug("Hello Debug")
	a.Log.Info("Hello Info")
	a.Log.Warn("Hello Warn")
	a.Log.Error("Hello Error")
}
