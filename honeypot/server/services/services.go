package services

import (
	"github.com/zhangweijie11/zSec/honeypot/server/services/mysql_server"
	"github.com/zhangweijie11/zSec/honeypot/server/services/redis_server"
	"github.com/zhangweijie11/zSec/honeypot/server/services/ssh_server"
	"github.com/zhangweijie11/zSec/honeypot/server/services/web"
	"github.com/zhangweijie11/zSec/honeypot/server/vars"
	"sync"
	"time"
)

type (
	HoneypotServices func(string, bool) error

	ServiceInfo struct {
		ServerName       string           `json:"server_name"`
		ListenAddr       string           `json:"listen_addr"`
		Flag             bool             `json:"flag"`
		HoneypotServices HoneypotServices `json:"honeypot_services"`
	}
)

var (
	Services   []ServiceInfo
	fnServices = map[string]HoneypotServices{"ssh": ssh_server.StartSsh, "redis": redis_server.StartRedis, "mysql": mysql_server.StartMysql, "web": web.StartWeb}
)

func init() {
	Services = make([]ServiceInfo, 0)
	for service, item := range vars.Config.Services {
		Services = append(Services, ServiceInfo{ServerName: service, ListenAddr: item.Addr, HoneypotServices: fnServices[service], Flag: item.Flag})
	}
}

func Start() {
	var wg sync.WaitGroup
	for _, s := range Services {
		wg.Add(1)
		go func(service ServiceInfo) {
			err := service.HoneypotServices(service.ListenAddr, service.Flag)
			_ = err
			wg.Done()
		}(s)
	}
	wg.Wait()

	for {
		time.Sleep(100 * time.Second)
	}
}
