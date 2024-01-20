package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zhangweijie11/zSec/honeypot/server/util"
	"github.com/zhangweijie11/zSec/honeypot/server/vars"
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
	timeout := time.Duration(1 * time.Second)
	client := http.Client{Timeout: timeout}

	return &HttpHook{HttpClient: client}, nil
}

func (hook *HttpHook) Fire(entry *logrus.Entry) (err error) {
	field := entry.Data

	data, ok := field["api"]
	fmt.Printf("data: %v, ok: %v\n", data, ok)
	if ok {
		urlApi := fmt.Sprintf("%v%v", vars.Config.Api.Addr, field["api"])
		data := entry.Message
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		secureKey := util.MakeSign(time.Now().Format("2006-01-02 15:04:05"), vars.Config.Api.Key)
		resp, err := hook.HttpClient.PostForm(urlApi, url.Values{"timestamp": {timestamp},
			"secureKey": {secureKey}, "data": {data}})

		fmt.Printf("resp: %v, err: %v\n", resp, err)
		if err != nil {
			return err
		}
	}

	return err
}

func (hook *HttpHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
