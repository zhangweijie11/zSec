package sniffer

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/zhangweijie11/zSec/honeypot/agent/logger"
	"github.com/zhangweijie11/zSec/honeypot/agent/settings"
	"net/url"
	"strings"
	"time"
)

var (
	device      string
	snapshotLen int32 = 1024
	promiscuous bool
	err         error
	handle      *pcap.Handle

	filter  = ""
	timeout = time.Duration(3)

	ApiUrl    string
	SecureKey string

	Ips []string

	ApiIp     string
	SensorIps []string
)

func init() {
	device = settings.InterfaceName
	ApiUrl = settings.ManagerUrl
	SecureKey = settings.SecKey

	Ips, err = GetIpList(device)

	urlParsed, err := url.Parse(ApiUrl)
	if err == nil {
		apiHost := urlParsed.Host
		ApiIp, _ = Host2Ip(strings.Split(apiHost, ":")[0])
		SensorIps = Ips
	}

	logger.Log.Infof("local address: %v, apiIp: %v", SensorIps, ApiIp)

	// 给hookHttp添加hook
	hookHttp, err := logger.NewHttpHook()
	if err == nil {
		logger.LogReport.Logger.AddHook(hookHttp)
	}
}

func Start() {
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer handle.Close()
	err = handle.SetBPFFilter(filter)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	ProcessPacket(packetSource.Packets())
}
