package logging

import (
	"github.com/sirupsen/logrus"
)

var LOG_LEVELS = map[string]logrus.Level{
	"DEBUG": logrus.DebugLevel,
	"INFO":  logrus.InfoLevel,
	"WARN":  logrus.WarnLevel,
	"ERROR": logrus.ErrorLevel,
	"FATAL": logrus.FatalLevel,
}

func CreateLogger(level string) *logrus.Entry {
	logrus.SetLevel(LOG_LEVELS[level])
	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})
	return contextLogger
}
