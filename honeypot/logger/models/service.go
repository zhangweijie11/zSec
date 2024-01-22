package models

import (
	"context"
	"time"
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

func (p *HoneypotMessage) Insert() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := CollectionService.InsertOne(ctx, p)
	return err
}
