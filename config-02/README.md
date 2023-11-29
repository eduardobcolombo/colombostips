# Config-02
## Importing configs to a package

This config tips shows how to use `envconfig` package from Kelsey Hightower, extended to import configs to a package.

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

As we did define the split_words to true, the applicatin will recognize my env variable splitted by underscore, like NEW_RELIC_LICENSE_KEY.

So if you run this app with the above command, the output should be: 

```
Using Log Level: info

Using NewRelic 
AppName: colombostips
LicenseKey: REDACTED

```

But you can see that it is printing the environment variable set.

