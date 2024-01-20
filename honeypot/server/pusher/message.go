package pusher

import (
	"encoding/json"
	"github.com/zhangweijie11/zSec/honeypot/server/logger"
)

type (
	HoneypotMessage struct {
		Timestamp   int64
		RawIp       string
		ProxyAddr   string
		ServiceType string
		User        string
		Password    string
		Data        map[string]interface{}
	}
)

func (m *HoneypotMessage) Build() (string, error) {
	ret, err := json.Marshal(m)
	return string(ret), err
}

func (m *HoneypotMessage) Send() error {
	message, err := m.Build()
	if err != nil {
		return err
	}

	go logger.LogHttp.WithField("api", "/api/service/").Info(message)
	return err
}
