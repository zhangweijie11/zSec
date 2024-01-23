package routers

import (
	"encoding/json"
	"github.com/zhangweijie11/zSec/proxy_honeypot/server/log"
	"github.com/zhangweijie11/zSec/proxy_honeypot/server/models"
	"github.com/zhangweijie11/zSec/proxy_honeypot/server/settings"
	"github.com/zhangweijie11/zSec/proxy_honeypot/server/util"
	"gopkg.in/macaron.v1"
	"strings"
)

func Index(ctx *macaron.Context) {
	_ = ctx.Req.ParseForm()
	log.Logger.Info(ctx.Req.Form)
	_, _ = ctx.Write([]byte("test only"))
}

func RecvData(ctx *macaron.Context) {
	_ = ctx.Req.ParseForm()
	timestamp := ctx.Req.Form.Get("timestamp")
	secureKey := ctx.Req.Form.Get("secureKey")
	data := ctx.Req.Form.Get("data")
	agentHost := ctx.Req.Form.Get("hostname")

	headers := ctx.Req.Header

	// get remote ips
	realIp := headers["X-Forwarded-For"]
	ips := make([]string, 0)
	if len(realIp) > 0 {
		t := strings.Split(realIp[0], ",")
		for _, ip := range t {
			sliceIp := strings.Split(ip, ".")
			if len(sliceIp) == 4 {
				ips = append(ips, strings.TrimSpace(ip))
			}

		}
	} else {
		ips = append(ips, ctx.Req.RemoteAddr)
	}

	mySecretKey := util.MakeSign(timestamp, settings.SECRET)
	if secureKey == mySecretKey {
		var h models.HttpRecord
		err := json.Unmarshal([]byte(data), &h)
		// log.Logger.Info(resp, err)
		agentIp := util.Address2Ip(ctx.Req.RemoteAddr)
		if err == nil {
			if len(ips) > 0 {
				agentIp = ips[0]
			}
			record := models.NewRecord(agentIp, agentHost, h)
			err = record.Insert()
			log.Logger.Infof("record: %v, err: %v", record, err)
		}
	} else {
		_, _ = ctx.Write([]byte("error"))
	}
}
