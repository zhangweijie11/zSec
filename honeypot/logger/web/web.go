package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangweijie11/zSec/honeypot/logger/settings"
	"github.com/zhangweijie11/zSec/honeypot/logger/web/routers"
)

func StartWeb() error {
	router := gin.Default()
	router.Use(gin.Logger())

	if settings.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	api := router.Group("/api/")
	{
		api.POST("/service/", routers.ServiceHandle)
		api.POST("/packet/", routers.PacketHandle)
	}

	err := router.Run(fmt.Sprintf("0.0.0.0:%v", settings.HttpPort))

	return err
}
