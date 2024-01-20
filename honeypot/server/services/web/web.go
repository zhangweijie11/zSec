package web

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangweijie11/zSec/honeypot/server/services/web/routers"
)

func Flagger(flag bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("flag", flag)
		ctx.Next()
	}
}

func StartWeb(addr string, flag bool) error {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(Flagger(flag))

	r.Any("/", routers.IndexHandle)
	err := r.Run(addr)
	return err
}
