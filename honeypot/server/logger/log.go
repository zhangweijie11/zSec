package logger

import "github.com/sirupsen/logrus"

var (
	Log     *logrus.Entry
	LogHttp *logrus.Entry
)

func init() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Log = logger.WithFields(logrus.Fields{"prefix": "honeypot-server"})

	LogHttp = logger.WithFields(logrus.Fields{"prefix": "honeypot-server", "report": true})
	httpHook, err := NewHttpHook()
	if err == nil {
		LogHttp.Logger.AddHook(httpHook)
	}
}
