package logtest

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var baseLogger = logrus.New()

type MyLogger struct {
	*logrus.Logger
}

func init() {
	baseLogger.SetOutput(os.Stdout)
	formatter := &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	baseLogger.SetFormatter(formatter)
	baseLogger.SetLevel(logrus.DebugLevel)
}

func Info(args ...interface{}) {
	message := fmt.Sprint(args...)
	baseLogger.Info(fmt.Sprintf("%s %s", "", message))
}

func Warn(args ...interface{}) {
	message := fmt.Sprint(args...)
	baseLogger.Warn(fmt.Sprintf("%s %s", "⚠️", message))
}

func Error(args ...interface{}) {
	message := fmt.Sprint(args...)
	baseLogger.Error(fmt.Sprintf("%s %s", "❌️", message))
}

func Fatal(args ...interface{}) {
	message := fmt.Sprint(args...)
	baseLogger.Fatal(fmt.Sprintf("%s %s", "☠️", message))
	baseLogger.Exit(1)
}
