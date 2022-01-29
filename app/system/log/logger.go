package log

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logOnce sync.Once
	logInst *logrus.Logger
)

func GetLogger() *logrus.Logger {
	logOnce.Do(func() {
		logInst = logrus.New()
	})
	return logInst
}

func Error(err error) {
	GetLogger().Errorf("Error: %s", err.Error())
}
