package logger

var (
	lokiAddress string
	labels      map[string]string
	usingLoki   bool = false
	logToFile   bool = false
	logFilePath string
)

type Config struct {
	LokiAddress string
	Labels      map[string]string
	UsingLoki   bool
	LogToFile   bool
	LogFilePath string
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
	logToFile = config.LogToFile
	logFilePath = config.LogFilePath
}
