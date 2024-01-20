package vars

import (
	"github.com/zhangweijie11/zSec/honeypot/server/config"
	"log"
	"sync"
)

var (
	RawIps sync.Map
	Flag   = true
	err    error
	Config config.Config
)

func init() {
	Config, err = config.ReadConfig()
	if err != nil {
		log.Panic(err)
	}
}
