package sniffer

import (
	"github.com/google/gopacket/pcap"
	"github.com/zhangweijie11/zSec/honeypot/agent/models"
	"github.com/zhangweijie11/zSec/honeypot/agent/vars"
	"net"
)

func GetIpList(deviceName string) (ips []string, err error) {
	devices, err := pcap.FindAllDevs()
	if err == nil {
		for _, device := range devices {
			if device.Name == deviceName {
				for _, addr := range device.Addresses {
					if addr.IP.To4() != nil {
						ips = append(ips, addr.IP.To4().String())
					}
				}
			}
		}
	}

	return ips, err
}

func Host2Ip(host string) (ip string, err error) {
	addr, err := net.LookupHost(host)
	if len(addr) > 0 {
		ip = addr[0]
	}

	return ip, err
}

func SliceContainsString(slice []string, str string) bool {
	m := make(map[string]bool)
	for _, v := range slice {
		m[v] = true
	}
	_, ok := m[str]
	return ok
}

func IsInWhite(conn *models.ConnectionInfo) (result bool) {

	if SliceContainsString(vars.HoneypotPolicy.WhiteIps, conn.SrcIp) ||
		SliceContainsString(vars.HoneypotPolicy.WhiteIps, conn.DstIp) ||
		SliceContainsString(vars.HoneypotPolicy.WhitePorts, conn.SrcPort) ||
		SliceContainsString(vars.HoneypotPolicy.WhitePorts, conn.DstPort) {
		result = true
	}

	return result
}
