# Config-01

This config tip shows how to use `envconfig` package from Kelsey Hightower.

To run this configuration demo, you will need to set an environment variable or use the below command to run.

`LOG_LEVEL=info go run main.go`

As we did define the split_words to true, the application will recognize the env variable split by underscore, like LOG_LEVEL.

```go
type Config struct {
	LogLevel string `split_words:"true" default:"debug"`
}
```

Look that we set the default value for this variable to `debug` and in the running, it is set as `info`

So if you run this app with the above command, the output should be: 

```
❯ LOG_LEVEL=info go run main.go
LogLevel: info
```

If you run it without passing the LOG_LEVEL variable, it should return the default value.

```
❯ go run main.go
LogLevel: debug
```


# References:

- https://github.com/kelseyhightower/envconfig
