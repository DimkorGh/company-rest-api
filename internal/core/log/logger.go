package log

import (
	"company-rest-api/internal/core/config"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	cnf *config.Config
}

func NewLogger(cnf *config.Config) *Logger {
	return &Logger{
		cnf: cnf,
	}
}

func (lgr *Logger) Initialize() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)

	logLevel, err := strconv.Atoi(lgr.cnf.Logger.Level)
	if err != nil {
		log.Fatalf("Error while converting log level to int %s", err.Error())
	}

	log.SetLevel(logrus.Level(logLevel))
}
