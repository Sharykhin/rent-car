package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() error {
	serviceID := os.Getenv("SERVICE_ID")
	if serviceID == "" {
		return fmt.Errorf("SERVICE_ID environment variable is not defined")
	}

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return fmt.Errorf("failed to parse LOG_LEVEL env variable: %v", err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "timestamp",
		},
	})
	logrus.SetLevel(level)
	Log = logrus.WithField("service", serviceID).Logger

	return nil
}
