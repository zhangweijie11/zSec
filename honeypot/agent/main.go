package main

import (
	"github.com/zhangweijie11/zSec/honeypot/agent/logger"
	"github.com/zhangweijie11/zSec/honeypot/agent/modules"
	"github.com/zhangweijie11/zSec/honeypot/agent/proxy"
	"github.com/zhangweijie11/zSec/honeypot/agent/sniffer"
)

func main() {
	_, err := modules.LoadPolicy()
	// logger.Log.Infof("load policy: %v, err: %v", policy, err)
	if err != nil {
		logger.Log.Panicln(err)
	}

	go proxy.Proxy()

	sniffer.Start()
}
