package log

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (lgr *Logger) Initialize() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)

	logLevel, err := strconv.Atoi(os.Getenv("ERROR_LEVEL"))
	if err != nil {
		log.Fatalf("Error while converting log level to int %s", err.Error())
	}

	log.SetLevel(logrus.Level(logLevel))
}
