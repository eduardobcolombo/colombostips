package newrelic

import (
	"fmt"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type Config struct {
	AppName    string `conf:"default:appName"`
	LicenseKey string `conf:"default:LicenceKey"`
}

func New(cfg Config) (*newrelic.Application, error) {

	fmt.Printf("Using NewRelic \nAppName: %s\nLicenseKey: %s \n\n", cfg.AppName, "REDACTED")

	return newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.AppName),
		newrelic.ConfigLicense(cfg.LicenseKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
}

func WrapHandleFunc(app *newrelic.Application, pattern string, handler func(http.ResponseWriter, *http.Request), options ...newrelic.TraceOption) (string, func(http.ResponseWriter, *http.Request)) {
	p, h := newrelic.WrapHandleFunc(app, pattern, handler)

	return p, h
}

func StartTransaction(nrApp *newrelic.Application, name string) *newrelic.Transaction {
	return nrApp.StartTransaction("Colombostips")
}
