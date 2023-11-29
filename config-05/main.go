package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ardanlabs/conf/v3"
	myLog "github.com/eduardobcolombo/colombostips/config-05/pkg/log"
	"github.com/eduardobcolombo/colombostips/config-05/pkg/newrelic"

	"go.uber.org/zap"
)

// Config represents a configuration type to receive env variables
type Config struct {
	Log      myLog.Config
	NewRelic newrelic.Config
	Port     string `conf:"default:8000"`
}

type App struct {
	log *zap.Logger
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

	const prefix = ""
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return 1, nil
		}
		return 1, fmt.Errorf("parsing config: %w", err)
	}

	logger := myLog.New(cfg.Log)
	defer logger.Sync()

	nrApp, err := newrelic.New(cfg.NewRelic)
	if err != nil {
		log.Fatal(err)
		return 1, err
	}

	app := App{log: logger}

	txn := newrelic.StartTransaction(nrApp, "Example transaction")
	defer txn.End()

	txn.AddAttribute("traceID", txn.GetTraceMetadata().TraceID)

	http.HandleFunc(newrelic.WrapHandleFunc(nrApp, "/", app.hello))
	port := fmt.Sprintf(":%s", cfg.Port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}

	return 0, nil
}

func (a App) hello(w http.ResponseWriter, req *http.Request) {
	// Write in the response
	fmt.Fprintf(w, "hello\n")
	// Write in the log
	a.log.Debug("Hello Debug")
	a.log.Info("Hello Info")
	a.log.Warn("Hello Warn")
	a.log.Error("Hello Error")
}
