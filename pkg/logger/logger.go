package logger

import (
	"github.com/involvestecnologia/mole/models"
	"github.com/sirupsen/logrus"
	graylog "gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

type logger struct {
	logger *logrus.Logger
}

func New(config models.ReadConfig) Log {

	logrus := logrus.New()
	logrus.AddHook(graylog.NewGraylogHook(config.Logstash.URL, map[string]interface{}{
		"application": config.AppName,
	}))

	return &logger{
		logger: logrus,
	}
}

//Error logs something to stdout as Error
func (l logger) Error(msg string) {
	l.logger.Error(msg)
}
