package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangweijie11/zSec/honeypot/logger/models"
	"github.com/zhangweijie11/zSec/honeypot/logger/settings"
	"github.com/zhangweijie11/zSec/honeypot/logger/util"
	"log"
)

func PacketHandle(ctx *gin.Context) {
	timestamp := ctx.PostForm("timestamp")
	secureKey := ctx.PostForm("secureKey")
	data := ctx.PostForm("data")
	remoteAddr := ctx.Request.RemoteAddr
	mySecureKey := util.MD5(fmt.Sprintf("%v%v", timestamp, settings.KEY))
	packetInfo := models.PacketInfo{}
	log.Printf("data: %v\n", data)
	if mySecureKey == secureKey {
		err := json.Unmarshal([]byte(data), &packetInfo)
		if err == nil {
			err := packetInfo.Insert()
			fmt.Printf("remoteAddr: %v, packetInfo: %v, err: %v\n", remoteAddr, packetInfo, err)
		}

		ctx.JSON(200, "ok")
	} else {
		ctx.JSON(200, "err")
	}
}

func ServiceHandle(ctx *gin.Context) {
	timestamp := ctx.PostForm("timestamp")
	secureKey := ctx.PostForm("secureKey")
	data := ctx.PostForm("data")
	remoteAddr := ctx.Request.RemoteAddr
	mySecureKey := util.MD5(fmt.Sprintf("%v%v", timestamp, settings.KEY))

	log.Printf("data: %v\n", data)
	if secureKey == mySecureKey {
		var message models.HoneypotMessage
		err := json.Unmarshal([]byte(data), &message)
		if err == nil {
			err := message.Insert()
			fmt.Printf("remoteAddr: %v, message: %v, err: %v\n", remoteAddr, message, err)
		}
	}
}
