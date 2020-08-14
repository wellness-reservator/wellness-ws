package config

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Datasource   string `required:"true"`
	LogLevel     string `envconfig:"LOG_LEVEL" default:"info"`
	MaxIdleConns int    `envconfig:"MAX_IDLE_CONNS" default:"2"`
	MaxOpenConns int    `envconfig:"MAX_OPEN_CONNS" default:"0"`
	DbTrace      bool   `envconfig:"DB_TRACE" default:"false"`
	ServeHttp    bool   `envconfig:"SERVE_HTTP" default:"false`
}

var ConfigLoaded Config

func LoadConfig() {
	err := envconfig.Process("wellness", &ConfigLoaded)
	logrus.Info("wellness config", ConfigLoaded)
	if err != nil {
		logrus.Panic(err.Error())
	}
}
