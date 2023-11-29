package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	myLog "github.com/eduardobcolombo/colombostips/config-04/pkg/log"
	"github.com/eduardobcolombo/colombostips/config-04/pkg/newrelic"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

// Config represents a configuration type to receive env variables
type Config struct {
	Log      myLog.Config
	NewRelic newrelic.Config `envconfig:"NEW_RELIC" desc:"NewRelic config"`
	Port     string          `split_words:"true" default:"8000"`
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

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("error processing envconfig %v", err)
		return 1, err
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
