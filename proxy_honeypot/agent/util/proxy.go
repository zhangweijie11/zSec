package util

import (
	"fmt"
	"github.com/elazarl/goproxy"
	goproxy_html "github.com/elazarl/goproxy/ext/html"
	"github.com/urfave/cli"
	"github.com/zhangweijie11/zSec/proxy_honeypot/agent/log"
	"github.com/zhangweijie11/zSec/proxy_honeypot/agent/modules"
	"github.com/zhangweijie11/zSec/proxy_honeypot/agent/settings"
	"github.com/zhangweijie11/zSec/proxy_honeypot/agent/vars"
	"net/http"
)

func init() {
	cfg := settings.Cfg
	sec := cfg.Section("proxy")
	vars.ProxyHost = sec.Key("HOST").MustString("")
	vars.ProxyPort = sec.Key("PORT").MustInt(1080)
	vars.DebugMode = sec.Key("DEBUG").MustBool(false)

}

func Start(ctx *cli.Context) {
	if ctx.IsSet("debug") {
		vars.DebugMode = ctx.Bool("debug")
	}

	if ctx.IsSet("port") {
		vars.ProxyPort = ctx.Int("port")
	}

	err := SetCA()
	log.Logger.Infof("caKey: %v, caCert: %v, set ca err: %v", vars.CaKey, vars.CaCert, err)

	proxy := goproxy.NewProxyHttpServer()
	log.Logger.Infof("proxy Start success, Listening on %v:%v ", vars.ProxyHost, vars.ProxyPort)
	// 指定Connect 请求方式为AlwaysMitm，启用 MIMT 功能
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxy.OnRequest().DoFunc(modules.ReqHandlerFunc)
	proxy.OnResponse(goproxy_html.IsWebRelatedText).DoFunc(modules.RespHandlerFunc)

	proxy.Verbose = vars.DebugMode

	log.Logger.Info(http.ListenAndServe(fmt.Sprintf("%v:%v", vars.ProxyHost, vars.ProxyPort), proxy))
}
