package ssh_server

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"github.com/zhangweijie11/zSec/honeypot/server/logger"
	"github.com/zhangweijie11/zSec/honeypot/server/pusher"
	"github.com/zhangweijie11/zSec/honeypot/server/util"
)

func StartSsh(addr string, flag bool) error {
	ssh.Handle(func(s ssh.Session) {
		// 送佛送到西，通过ssh蜜罐再指引黑客去下一个蜜罐打卡。
		s.Write([]byte(fmt.Sprintf("您的来源IP:%v不在可信列表范围内，"+
			"按公司的安全规范，请先登录跳板机（jumper.sec.lu），再用跳板机登录服务器。\n", s.RemoteAddr())))
	})

	passwordOpt := ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		result := false

		if ctx.User() == "root" && password == "123456" {
			result = true
		}

		if flag {
			localAddr := ctx.LocalAddr().String()
			remoteAddr := ctx.RemoteAddr().String()
			rawIp, ProxyAddr, timeStamp := util.GetRawIp(remoteAddr, localAddr)
			logger.Log.Warningf("timestamp: %v, rawIp: %v, proxyAddr: %v, user: %v, password: %v",
				timeStamp, rawIp, ProxyAddr, ctx.User(), password)

			var message pusher.HoneypotMessage
			message.Timestamp = timeStamp
			message.RawIp = rawIp
			message.ProxyAddr = ProxyAddr.String()
			message.User = ctx.User()
			message.Password = password

			strMessage, _ := message.Build()
			logger.Log.Info(strMessage)
			_ = message.Send()
		}

		return result
	})

	logger.Log.Warningf("start ssh service on %v", addr)
	err := ssh.ListenAndServe(addr, nil, passwordOpt)
	return err
}
