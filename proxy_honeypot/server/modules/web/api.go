package web

import (
	"fmt"
	"github.com/zhangweijie11/zSec/honeypot/logger/settings"
	"github.com/zhangweijie11/zSec/proxy_honeypot/server/log"
	"github.com/zhangweijie11/zSec/proxy_honeypot/server/modules/web/routers"
	"gopkg.in/macaron.v1"
	"net/http"
)

func Start() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/", routers.Index)
	m.Post("/api/send", routers.RecvData)
	log.Logger.Infof("start web server at: %v", settings.HttpPort)
	log.Logger.Debug(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", settings.HttpPort), m))
}
