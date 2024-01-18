package sensor

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/zhangweijie11/zSec/traffic_analysis/sensor/misc"
	"github.com/zhangweijie11/zSec/traffic_analysis/sensor/settings"
	"time"
)

var (
	device      string
	snapshotLen int32 = 1024
	promiscuous bool  = true
	err         error
	timeout     time.Duration = pcap.BlockForever
	handle      *pcap.Handle

	DebugMode bool
	filter    = ""

	ApiUrl    string
	SecureKey string
)

func init() {
	device = settings.DeviceName
	DebugMode = settings.DebugMode
	filter = settings.FilterRule

	cfg := settings.Cfg
	sec := cfg.Section("server")
	ApiUrl = sec.Key("API_URL").MustString("")
	SecureKey = sec.Key("API_KEY").MustString("")

}

func Start(ctx *cli.Context) {
	if ctx.IsSet("debug") {
		DebugMode = ctx.Bool("debug")
	}
	if DebugMode {
		misc.Log.Logger.Level = logrus.DebugLevel
	}

	if ctx.IsSet("length") {
		snapshotLen = int32(ctx.Int("len"))
	}
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		misc.Log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	if ctx.IsSet("filter") {
		filter = ctx.String("filter")
	}

	err := handle.SetBPFFilter(filter)
	misc.Log.Infof("set SetBPFFilter: %v, err: %v", filter, err)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	ProcessPackets(packetSource.Packets())
}
