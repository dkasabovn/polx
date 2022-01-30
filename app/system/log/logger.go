package log

import (
	"os"
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
		logInst.SetOutput(os.Stdout)
	})
	return logInst
}

func Error(err error) {
	GetLogger().Errorf("Error: %s", err.Error())
}

func InfoStruct(data interface{}) {
	GetLogger().Infof("%+v", data)
}
