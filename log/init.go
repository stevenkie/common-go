package commonlog

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	formatter := new(prefixed.TextFormatter)
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	SetFormatter(formatter)
}

// Config used for logrus config
type Config struct {
	LogLevel  string
	ShortPath bool
	Formatter logrus.Formatter
}

// SetLogConfig set config for logrus
func SetLogConfig(config Config) {
	if len(config.LogLevel) > 0 {
		baseLogger.SetLevel(config.LogLevel)
	}

	if config.Formatter != nil {
		SetFormatter(config.Formatter)
	}

	shortPathFile = config.ShortPath
}
