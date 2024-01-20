package util

import (
	"fmt"
	"github.com/zhangweijie11/zSec/honeypot/server/vars"
	"net"
	"strconv"
	"strings"
	"time"
)

func GetRawIpByConn(conn net.Conn) (string, net.TCPAddr, int64) {
	remoteAddr := conn.RemoteAddr().String()
	localAddr := conn.LocalAddr().String()
	return GetRawIp(remoteAddr, localAddr)
}

func GetRawIp(remoteAddr, localAddr string) (string, net.TCPAddr, int64) {
	var (
		rawIp     string
		ProxyAddr net.TCPAddr
		timeStamp int64
	)

	k := fmt.Sprintf("%v_%v", remoteAddr, localAddr)
	v, ok := vars.RawIps.Load(k)
	fmt.Printf("k: %v, v: %v, ok: %v\n", k, v, ok)
	if ok {
		value, ok := v.(string)
		if ok {
			t := strings.Split(value, "_")
			if len(t) == 2 {
				rawIp = t[1]
				ProxyAddrStr := t[0]
				tt := strings.Split(ProxyAddrStr, "@")
				if len(tt) == 2 {

					timeStamp, _ = strconv.ParseInt(tt[0], 10, 64)
					proxyIpPort := tt[1]
					ttt := strings.Split(proxyIpPort, ":")
					// fmt.Printf("ttt: %v, len(ttt): %v\n", ttt, len(ttt))
					if len(ttt) == 2 {
						ProxyAddr.IP = StrToIp(ttt[0])
						port, _ := strconv.Atoi(ttt[1])
						ProxyAddr.Port = port
					}
				}
			}
		}
	}

	return rawIp, ProxyAddr, timeStamp
}

func StrToIp(ip string) net.IP {
	return net.ParseIP(ip)
}

func DelExpireIps(timeoutSec int64) {
	vars.RawIps.Range(func(key, value interface{}) bool {
		v, ok := value.(string)
		if ok {
			timestamp := getTimestamp(v)
			if time.Now().Unix()-timestamp >= timeoutSec {
				vars.RawIps.Delete(key)
			}
		}
		return ok
	})
}

func getTimestamp(v string) int64 {
	var timestamp int64
	t := strings.Split(v, "_")
	if len(t) == 2 {
		ProxyAddrStr := t[0]
		tt := strings.Split(ProxyAddrStr, "@")
		if len(tt) == 2 {
			timestamp, _ = strconv.ParseInt(tt[0], 10, 64)
		}
	}
	return timestamp
}
