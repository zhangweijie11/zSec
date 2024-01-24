package routers

import (
	"github.com/go-macaron/session"
	"github.com/zhangweijie11/zSec/proxy_honeypot/manager/models"
	"gopkg.in/macaron.v1"
)

func Dash(ctx *macaron.Context, sess session.Store) {
	if sess.Get("admin") != nil {
		totalRecord, totalPassword, err := models.DashTotal()
		passwords, err := models.DashPassword()
		urls, err := models.DashUrls()
		evilIps, err := models.DashIps()
		_ = err

		ctx.Data["total_record"] = totalRecord
		ctx.Data["total_password"] = totalPassword

		ctx.Data["passwords"] = passwords
		ctx.Data["urls"] = urls
		ctx.Data["evil_ips"] = evilIps

		ctx.HTML(200, "dash")
	} else {
		ctx.Redirect("/admin/login/")
	}
}
