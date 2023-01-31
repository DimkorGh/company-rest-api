package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (logger *Logger) Initialize() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)

	logLevel, err := strconv.Atoi(os.Getenv("ERROR_LEVEL"))
	if err != nil {
		log.Fatalf("Error while converting log level to int %s", err.Error())
	}

	log.SetLevel(logrus.Level(logLevel))
}
