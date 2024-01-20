package main

import (
	"github.com/zhangweijie11/zSec/honeypot/server/proxy"
	"github.com/zhangweijie11/zSec/honeypot/server/services"
)

func main() {
	go services.Start()
	proxy.StartProxy()
}
