package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zhangweijie11/zSec/honeypot/agent/settings"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/util"
	"net/http"
	"net/url"
	"time"
)

type (
	HttpHook struct {
		HttpClient http.Client
	}
)

func NewHttpHook() (*HttpHook, error) {
	timeout := 1 * time.Second
	client := http.Client{Timeout: timeout}

	return &HttpHook{HttpClient: client}, nil
}

func (hook *HttpHook) Fire(entry *logrus.Entry) (err error) {
	var serverUrl string
	field := entry.Data
	serverUrl = settings.ManagerUrl

	_, ok := field["api"]
	if ok {
		urlApi := fmt.Sprintf("%v%v", serverUrl, field["api"])
		data := entry.Message
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		secureKey := util.MakeSign(time.Now().Format("2006-01-02 15:04:05"), settings.SecKey)
		resp, err := hook.HttpClient.PostForm(urlApi, url.Values{"timestamp": {timestamp}, "secureKey": {secureKey}, "data": {data}})
		fmt.Printf("resp: %v, err: %v\n", resp, err)

	}

	return err
}

func (hook *HttpHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
