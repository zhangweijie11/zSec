package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangweijie11/zSec/honeypot/server/logger"
	"github.com/zhangweijie11/zSec/honeypot/server/pusher"
	"github.com/zhangweijie11/zSec/honeypot/server/util"
	"net/http"
)

var httpAddr = "127.0.0.1:8000"

func IndexHandle(ctx *gin.Context) {
	_, ok := ctx.Get("flag")
	_ = ctx.Request.ParseForm()
	params := ctx.Request.Form

	remoteAddr := ctx.Request.RemoteAddr
	host := ctx.Request.Host

	body := make([]byte, 0)
	n, err := ctx.Request.Body.Read(body)
	logger.Log.Infof("n: %v, err: %v", n, err)
	if ok {
		rawIp, ProxyAddr, timeStamp := util.GetRawIp(remoteAddr, httpAddr)
		logger.Log.Warnf("rawIp: %v, proxyAddr: %v, timestamp: %v", rawIp, ProxyAddr, timeStamp)
		var message pusher.HoneypotMessage
		message.Timestamp = timeStamp
		message.RawIp = rawIp
		message.ProxyAddr = ProxyAddr.String()

		data := make(map[string]interface{})
		data["body"] = body

		message.Data = data
		strMessage, _ := message.Build()
		logger.Log.Info(strMessage)
		_ = message.Send()
	}

	ctx.String(http.StatusOK, fmt.Sprintf("Hello, World! \nremote_addr: %v, host: %v, param: %v, body: %v\n",
		remoteAddr, host, params, string(body)))
}
