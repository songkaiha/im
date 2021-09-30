package models

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type LoggerType struct {
	File     *os.File
	FileName string
	Do       *logrus.Logger
}

var (
	Logger      *LoggerType
	loggerError error
)

func init() {
	Logger = &LoggerType{}
	Logger.Do = logrus.New()
	Logger.FileName = "logs/MemStats" + time.Now().Format("20060102") + ".log"
	Logger.File, loggerError = os.OpenFile(Logger.FileName, os.O_CREATE|os.O_WRONLY, 0666)
	if loggerError != nil {
		panic(loggerError)
	}
	Logger.Do.Out = Logger.File
	Logger.Do.Level = logrus.InfoLevel
	Logger.Do.Info("start-logger-do")
	//Logger.Do.Warn("warning do")

}

func GetLogger(title string) *logrus.Logger {
	fileName := "logs/" + title + time.Now().Format("20060102") + ".log"
	if Logger.FileName == fileName {
		return Logger.Do
	}
	Logger.FileName = fileName
	Logger.File, loggerError = os.OpenFile(Logger.FileName, os.O_CREATE|os.O_WRONLY, 0666)
	if loggerError != nil {
		return Logger.Do
	}
	Logger.Do.Out = Logger.File
	Logger.Do.Info("start-logger-new-day-do")
	return Logger.Do
}
