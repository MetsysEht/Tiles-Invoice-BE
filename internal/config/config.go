package config

// App contains application-specific config values

type Config struct {
	App App
}

type App struct {
	Env             string
	ServiceName     string
	Hostname        string
	Port            string
	ShutdownTimeout int
	ShutdownDelay   int
	GitCommitHash   string
	Interfaces      struct {
		Service NetworkInterfaces
	}
}

type NetworkInterfaces struct {
	GrpcServerAddress     string
	HttpServerAddress     string
	InternalServerAddress string
}
