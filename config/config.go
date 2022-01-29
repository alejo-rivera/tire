package config

type LoggingType string

const (
	Release = LoggingType("release")
	Debug   = LoggingType("debug")
)

type Data struct {
	ServerConfig *GraphServerConfig `env:",prefix=GRAPH_SERVER_"`
	LoggingLevel LoggingType        `env:"LOGGING_LEVEL,default=debug"`
}

type GraphServerConfig struct {
	Port   int  `env:"PORT,default=8181"`
	Enable bool `env:"ENABLE,default=false"`
}
