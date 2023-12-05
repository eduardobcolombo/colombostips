# Config-03

## Importing configs to a package using docker with prefix

This config tip shows how to use the `conf/v3` package from [ArdanLabs](https://github.com/ardanlabs/conf), which is extended to import configs to a package.

To run this configuration demo, you must set the environment variables in the `config.env` file, feel free to copy `config.env.example`. The content should looks like the below.

```
COLOMBOSTIPS_NEW_RELIC_LICENSE_KEY=some_valid_licence_key
COLOMBOSTIPS_NEW_RELIC_APP_NAME=colombostips
COLOMBOSTIPS_LOG_LEVEL=error
```

Note that we are using the prefix `COLOMBOSTIPS` for all env variables.
 
The run the docker commands to spin up the application.

```
docker compose up -d --build

docker compose logs
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
	var logCfg = logger.Config{
		Level: cfg.Log.Level,
	}

	logger := logger.New(logCfg)
...
```

In `pkg/newrelic/newrelic.go`, we defined another config struct which will be explicitly filled in the main.go. In this case, you are not hiding the configuration but setting it explicitly.

```go
... //main.go
	var nrCfg = newrelic.Config{
		AppName:    cfg.NewRelic.AppName,
		LicenseKey: cfg.NewRelic.LicenseKey,
	}

	// skipping the newrelic return for this demo
	if _, err = newrelic.New(nrCfg); err != nil {
		return fmt.Errorf("starting newrelic: %w", err)
	}
...
```

You can see that it is printing the environment variable set.

Note: Once this app only prints a output and did not hang like a http webserver, the docker container will stop right after.

So if you run this app with the above command, the output should be: 

```
config-03-go_config-1  | Using Log Level: info
config-03-go_config-1  | Using NewRelic 
config-03-go_config-1  | AppName: colombostips
config-03-go_config-1  | LicenseKey: REDACTED
config-03-go_config-1  | 
config-03-go_config-1  | 2023/12/05 13:07:07 starting newrelic: license length is not 40

```

You can see that it is printing the environment variable set.

Note: Once this app only prints a output and did not hang like a http webserver, the docker container will stop right after.


# References:

https://github.com/ardanlabs/conf

