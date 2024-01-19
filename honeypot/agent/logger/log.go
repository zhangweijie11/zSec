package logger

import "github.com/sirupsen/logrus"

var (
	Log       *logrus.Entry
	LogReport *logrus.Entry
)

func init() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})

	logger.Level = logrus.DebugLevel
	Log = logger.WithFields(logrus.Fields{"prefix": "honeypot agent"})
	LogReport = logrus.New().WithFields(logrus.Fields{"prefix": "honeypot agent", "report": true})
}
