package web

import (
	"fmt"
	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"github.com/urfave/cli"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/util"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/web/routers"
	"gopkg.in/macaron.v1"
	"net/http"
)

func RunWeb(ctx *cli.Context) (err error) {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())
	m.Use(cache.Cacher())

	m.Get("/", routers.Index)
	m.Get("/http/", routers.HttpReq)
	m.Get("/dns/", routers.Dns)

	m.Post("/api/packet/", routers.SendPacket)
	m.Post("/api/http/", routers.SendHTML)
	m.Post("/api/dns/", routers.SendDns)

	if ctx.IsSet("host") {
		HTTP_HOST = ctx.String("host")
	}

	if ctx.IsSet("port") {
		HTTP_PORT = ctx.Int("port")
	}

	util.Log.Infof("run server on %v", fmt.Sprintf("%v:%v", HTTP_HOST, HTTP_PORT))
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", HTTP_HOST, HTTP_PORT), m)

	return err
}
