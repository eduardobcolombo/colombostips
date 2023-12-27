package newrelic

import (
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
)

// Config has all the needed values to create the NewRelic.
type Config struct {
	AppName    string
	LicenseKey string
}

// NewRelic has all that is needed to use NewRelic pkg.
type NewRelic struct {
	app *newrelic.Application
}

// New provides a NewRelic object as a Factory pattern.
func New(cfg Config) (*NewRelic, error) {
	fmt.Printf("Using NewRelic \nAppName: %s\nLicenseKey: %s \n\n", cfg.AppName,
		"REDACTED")

	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.AppName),
		newrelic.ConfigLicense(cfg.LicenseKey),
	)
	if err != nil {
		return nil, err
	}

	nr := NewRelic{
		app: nrApp,
	}

	return &nr, nil
}
