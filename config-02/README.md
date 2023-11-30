# Config-02
## How to import environment variables into a package.

This config tips shows how to use `conf/v3` package from [ArdanLabs](https://github.com/ardanlabs/conf), extended to import configs to a package.

To run this configuration demo, you will need to set the below environment variables or use the below command to run.
```
NEW_RELIC_LICENSE_KEY=some_valid_licence_key
NEW_RELIC_APP_NAME=colombostips
LOG_LEVEL=info
```
 
```
NEW_RELIC_LICENSE_KEY=some_valid_licence_key \
NEW_RELIC_APP_NAME=colombostips \
LOG_LEVEL=info \
go run main.go
```

In additional we are passing the NewRelic configuration. Notice that we did use a local config type to get environment variables and we are passing the values using the `newrelic.Config` type explicit to keep the code easy to understand.

```go
... //main.go
	cfg := struct {
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

In `pkg/log/log.go`, we defined another config struct which will be explicit filled in the main.go.
```go
... // main.go
	logLevel := myLog.New(myLog.Config{
		Level: cfg.Log.Level,
	})
...
```

In `pkg/newrelic/newrelic.go`, we defined another config struct which will be explicit filled in the main.go. In this case, you are not hidden the configuration, but set it explicitly.

```go
... //main.go
	_, err = newrelic.New(newrelic.Config{
		AppName:    cfg.NewRelic.AppName,
		LicenseKey: cfg.NewRelic.LicenseKey,
	})
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

