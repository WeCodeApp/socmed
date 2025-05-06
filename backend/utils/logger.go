package utils

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger() {
	logger = logrus.New()

	// Create logs directory if it doesn't exist
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		_ = os.Mkdir("logs", 0755)
	}

	// Log to both file and stdout
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		logger.SetOutput(os.Stdout)
	}

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func GetLogger() *logrus.Logger {
	return logger
}
