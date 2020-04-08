package logger

import (
	"github.com/involvestecnologia/mole/models"
	"github.com/sirupsen/logrus"
	graylog "gopkg.in/gemnasium/logrus-graylog-hook.v2"
)

func New(config models.ReadConfig) *logrus.Logger {

	log := logrus.New()
	log.AddHook(graylog.NewGraylogHook(config.Logstash.URL, map[string]interface{}{
		"application": config.AppName,
	}))

	return log
}
