# Config-04
## Importing configs to a package using kubernetes and zap for logging


This config tip shows how to use the `conf/v3` package from [ArdanLabs](https://github.com/ardanlabs/conf), which is extended to import configs to a package.

To run this configuration demo, you will need to set the below environment variables to the deploy/server.env. Feel free to use the server.env.example.

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
		Port string `conf:"default:8000"`
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

	nr, err := newrelic.New(nrCfg)
	if err != nil {
		return fmt.Errorf("initiating newrelic: %w", err)
	}
...
```

As we are running on top of `k8s` (kubernetes), we need a cluster. In order to make it easy, this repo provides a `make file` with some shortcuts to create your own cluster with `KinD` and some make commands to get the appication up and down.
Please explore the Makefile.

```
make kind-init
```

kind-init will: 
- create a cluster using kind;
- set the namespace in the context;
- build the image with the go app;
- load the image into k8s;
- apply kind ymls from ./k8s folder and create the kubernetes infrastructure;
- generate the secrets with your environment variables.

Note: Of course you should use some Vault techniques to store your secrets, but for this example it is set only in the kube secrets.

You can cleanup it using the below command.
```
make kind-clean
```

Alternatively you can use k9s to explore your pods, if you are not so familiar with kube ctl cmds.
If you inspect Kubernetes logs you should see the below:

```
LogLevel: info

Using NewRelic 
AppName: colombostips
LicenseKey: REDACTED

```

As kubernetes uses a private network you need to do a fast-forward to expose the port so you can access your app from your terminal.
You can try using make command.
```
make expose
```

```
curl localhost:8080
```

So if you hit the api with the above command, the output should be: 
```
Hello
```

Note that in the logs it will shows only the logs in the log level set. In this case the `debug` log is not shown in the logs.
```
{"level":"info","timestamp":"2023-11-29T19:30:25.465Z","caller":"src/main.go:75","msg":"Hello Info","pid":1}
{"level":"warn","timestamp":"2023-11-29T19:30:25.465Z","caller":"src/main.go:76","msg":"Hello Warn","pid":1}
{"level":"error","timestamp":"2023-11-29T19:30:25.465Z","caller":"src/main.go:77","msg":"Hello Error","pid":1...
...
```

Note: 
If you change something in your go code, you must reload you app by running `make kind-reload`.
