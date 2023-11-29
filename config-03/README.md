# Config-03

## Importing configs to a package using docker with prefix

This config tips shows how to use `envconfig` package from Kelsey Hightower, extended to import configs to a package.

To run this configuration demo, you will need to set the below environment variables within the config.env file.
```
COLOMBOSTIPS_NEW_RELIC_LICENSE_KEY=some_valid_licence_key
COLOMBOSTIPS_NEW_RELIC_APP_NAME=colombostips
COLOMBOSTIPS_LOG_LEVEL=info
```

Note that we are using the prefix `COLOMBOSTIPS` for all env variables.
 
The run the docker commands to spin up the application.

```
docker compose up -d --build

docker compose logs
```

As we did define the split_words to true, the applicatin will recognize my env variable splitted by underscore, like LOG_LEVEL.

In additional we are passing the NewRelic configuration. Notice that we did use `envconfig:"NEW_RELIC"` and point it to `newrelic.Config` which is our newrelic package.

```go
type Config struct {
	Log      myLog.Config
	NewRelic newrelic.Config `envconfig:"NEW_RELIC" desc:"NewRelic config"`
}

```
In `pkg/log/log.go`, we defined another config struct which is the `myLog.Config` in the main.go
```go
type Config struct {
	Level string
}
```

In `pkg/newrelic/newrelic.go`, we defined another config struct which is the `newrelic.Config` in the main.go

```go
type Config struct {
	AppName    string `split_words:"true" desc:"application name"`
	LicenseKey string `split_words:"true" desc:"license key"`
}
```

So if you run this app with the above command, the output should be: 

```
config-03-go_config-1  | Using Log Level: info
config-03-go_config-1  | 
config-03-go_config-1  | Using NewRelic 
config-03-go_config-1  | AppName: colombostips
config-03-go_config-1  | LicenseKey: REDACTED
config-03-go_config-1  | 
config-03-go_config-1  | 
config-03-go_config-1 exited with code 0

```

You can see that it is printing the environment variable set.

Note: Once this app only prints a output and did not hang like a http webserver, the docker container will stop right after.
