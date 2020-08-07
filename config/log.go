package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitializeLog initialize log
func InitializeLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stdout)

	lvl, err := logrus.ParseLevel(ConfigLoaded.LogLevel)

	if err != nil {
		lvl = logrus.InfoLevel
	}

	// Only log the warning severity or above.
	logrus.SetLevel(lvl)
}
