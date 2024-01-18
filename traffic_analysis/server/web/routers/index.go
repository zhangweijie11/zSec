package routers

import (
	"context"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/models"
	"gopkg.in/macaron.v1"
)

func Index(ctx *macaron.Context) {
	ctxs := context.Background()
	info, _ := models.ListEvilInfo(ctxs)
	ctx.Data["info"] = info
	ctx.HTML(200, "index")
}

func HttpReq(ctx *macaron.Context) {
	info, _ := models.ListEvilHttpReq()
	ctx.Data["info"] = info
	ctx.HTML(200, "http_req")
}

func Dns(ctx *macaron.Context) {
	info, _ := models.ListEvilDns()
	ctx.Data["info"] = info
	ctx.HTML(200, "dns")
}
