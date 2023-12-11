# Config-02
## How to import environment variables into a package.

This config tip shows how to use the `conf/v3` package from [ArdanLabs](https://github.com/ardanlabs/conf), which is extended to import configs to a package.

To run this configuration demo, you must set the environment variables below or use the below command to run.

```
NEW_RELIC_LICENSE_KEY=some_valid_licence_key \
NEW_RELIC_APP_NAME=colombostips \
LOG_LEVEL=info \
go run main.go
```


Notice that we used a local config type to get environment variables and are passing the values using the `newrelic.Config` type explicitly to keep the code easy to understand. In addition to that, we are not exposing the main Config, avoiding passing it along to the application.

```go
... //main.go
	var cfg = struct {
		Log struct {
			Level string `conf:"default:error"`
		}
		NewRelic struct {
			AppName    string `conf:"default:appName"`
			LicenseKey string `conf:"default:LicenceKey"`
		}
	}{}
...
```

In `pkg/log/log.go`, we defined another config struct which will be explicitly filled in the main.go.
```go
... // main.go
	logCfg := logger.Config{
		Level: cfg.Log.Level,
	}

	logger := logger.New(logCfg)
...
```

In `pkg/newrelic/newrelic.go`, we defined another config struct which will be explicitly filled in the main.go. In this case, you are not hiding the configuration but setting it explicitly.

```go
... //main.go
	nrCfg := newrelic.Config{
		AppName:    cfg.NewRelic.AppName,
		LicenseKey: cfg.NewRelic.LicenseKey,
	}

	// skipping the newrelic return for this demo
	if _, err = newrelic.New(nrCfg); err != nil {
		return fmt.Errorf("starting newrelic: %w", err)
	}
...
```

So if you run this app with the above command, the output should be: 

```
Using Log Level: info

Using NewRelic 
AppName: colombostips
LicenseKey: REDACTED

```

But you can see that it is printing the environment variable set.

# References:

https://github.com/ardanlabs/conf
