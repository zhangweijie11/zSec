package redis_server

import (
	"github.com/zhangweijie11/zSec/honeypot/server/logger"
	"github.com/zhangweijie11/zSec/honeypot/server/services/redis_server/redis"
)

func StartRedis(addr string, flag bool) error {
	logger.Log.Warningf("start redis service on %v", addr)
	err := redis.Run(addr, flag)
	return err
}
