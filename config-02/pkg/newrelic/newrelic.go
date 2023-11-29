package newrelic

import (
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type Config struct {
	AppName    string `split_words:"true" desc:"application name"`
	LicenseKey string `split_words:"true" desc:"license key"`
}

func New(cfg Config) (*newrelic.Application, error) {

	fmt.Printf("Using NewRelic \nAppName: %s\nLicenseKey: %s \n\n", cfg.AppName,
		"REDACTED")

	return newrelic.NewApplication(
		newrelic.ConfigAppName(cfg.AppName),
		newrelic.ConfigLicense(cfg.LicenseKey),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
}
