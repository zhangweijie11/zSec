package routers

import (
	"encoding/json"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/audit"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/models"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/settings"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/util"
	"gopkg.in/macaron.v1"
)

func SendPacket(ctx *macaron.Context) {
	_ = ctx.Req.ParseForm()
	timestamp := ctx.Req.Form.Get("timestamp")
	secureKey := ctx.Req.Form.Get("secureKey")
	data := ctx.Req.Form.Get("data")
	sensorIp := ctx.Req.RemoteAddr

	if secureKey == util.MakeSign(timestamp, settings.SECRET) {
		var packet models.ConnectionInfo
		err := json.Unmarshal([]byte(data), &packet)
		// util.Log.Errorf("err: %v, packet: %v", err, packet)
		if err == nil {
			go func() {
				_, _, _ = audit.PacketAduit(sensorIp, packet)
			}()
		}
	}
}

func SendHTML(ctx *macaron.Context) {
	_ = ctx.Req.ParseForm()
	timestamp := ctx.Req.Form.Get("timestamp")
	secureKey := ctx.Req.Form.Get("secureKey")
	data := ctx.Req.Form.Get("data")
	sensorIp := ctx.Req.RemoteAddr

	if secureKey == util.MakeSign(timestamp, settings.SECRET) {
		var req models.HttpReq
		err := json.Unmarshal([]byte(data), &req)
		// util.Log.Errorf("err: %v, req: %v", err, req)
		if err == nil {
			go func() {
				_, _, _ = audit.HttpAudit(sensorIp, req)
			}()
		}
	}
}

func SendDns(ctx *macaron.Context) {
	_ = ctx.Req.ParseForm()
	timestamp := ctx.Req.Form.Get("timestamp")
	secureKey := ctx.Req.Form.Get("secureKey")
	data := ctx.Req.Form.Get("data")
	sensorIp := ctx.Req.RemoteAddr

	if secureKey == util.MakeSign(timestamp, settings.SECRET) {
		var dns models.Dns
		err := json.Unmarshal([]byte(data), &dns)
		// util.Log.Errorf("err: %v, req: %v", err, req)
		if err == nil {
			go func() {
				_, _, _ = audit.DnsAudit(sensorIp, dns)
			}()
		}
	}
}
