package newrelic

import (
	"fmt"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// Config has all the needed values to create the NewRelic.
type Config struct {
	AppName    string
	LicenseKey string
}

// NewRelic has all that is needed to use NewRelic pkg.
type NewRelic struct {
	App *newrelic.Application
}

// New provides a NewRelic object as a Factory pattern.
func New(cfg Config) (NewRelic, error) {

	// hack for this demo
	fmt.Printf("Using NewRelic \nAppName: %s\nLicenseKey: %s \n\n", cfg.AppName,
		"REDACTED")

	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.AppName),
		newrelic.ConfigLicense(cfg.LicenseKey),
	)

	if err != nil {
		return NewRelic{}, err
	}

	return NewRelic{
		App: nrApp,
	}, nil
}

func WrapHandleFunc(app *newrelic.Application, pattern string, handler func(http.ResponseWriter, *http.Request), options ...newrelic.TraceOption) (string, func(http.ResponseWriter, *http.Request)) {
	p, h := newrelic.WrapHandleFunc(app, pattern, handler)

	return p, h
}

func StartTransaction(nrApp *newrelic.Application, name string) *newrelic.Transaction {
	return nrApp.StartTransaction("Colombostips")
}
