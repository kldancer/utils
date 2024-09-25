package logtest

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type IconFormatter struct {
	Icon string
}

func (f *IconFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	message := fmt.Sprintf("%s %s", entry.Message, f.Icon)
	entry.Message = message

	// 使用默认的TextFormatter来格式化日志
	formatter := &logrus.TextFormatter{DisableTimestamp: true}
	return formatter.Format(entry)
}
