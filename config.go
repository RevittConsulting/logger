package logger

var (
	lokiAddress string
	labels      map[string]string
	usingLoki   bool = false
)

type Config struct {
	LokiAddress string
	Labels      map[string]string
	UsingLoki   bool
}

// InitLoggerConfig initialises logger with config.
// Labels example: map[string]string{"app": My App, env: production}
func InitLoggerConfig(config *Config) {
	if len(config.LokiAddress) == 0 {
		panic("loki address is required")
	}
	lokiAddress = config.LokiAddress
	usingLoki = config.UsingLoki
	labels = config.Labels
}
