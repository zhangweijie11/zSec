package main

import (
	"bytes"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
	"time"
)

var (
	//device               = "en0"
	device               = "lo0"
	snapshotLength int32 = 1024
	promiscuous          = false // 混杂模式
	timeout              = 5 * time.Second

	handle *pcap.Handle
	err    error

	//filter   = "(tcp and dst port 21) or  (tcp and dst port 80) or (tcp and dst port 25) or (tcp and dst port 110) or (tcp and dst port 5601)"
	filter   = "tcp and dst port 5002"
	userList = []string{"user", "username", "login", "login_user", "manager", "user_name", "usr"}
	passList = []string{"pass", "password", "login_pass", "pwd", "passwd"}
)

func main() {
	handle, err = pcap.OpenLive(device, snapshotLength, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	// 设置过滤器
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		processPacket(packet)
	}
}

func processPacket(packet gopacket.Packet) {
	var (
		fromIp   string
		destIp   string
		srcPort  string
		destPort string
	)

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		fromIp = ip.SrcIP.String()
		destIp = ip.DstIP.String()
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		srcPort = tcp.SrcPort.String()
		destPort = tcp.DstPort.String()
	}

	applicationLayer := packet.ApplicationLayer()
	if applicationLayer != nil {
		payload := applicationLayer.Payload()
		if user, ok := checkUsername(payload); ok {
			_ = user
			fmt.Printf("%v:%v->%v:%v, %v\n", fromIp, srcPort, destIp, destPort, string(payload))
		}
		if pass, ok := checkPassword(payload); ok {
			_ = pass
			fmt.Printf("%v:%v->%v:%v, %v\n", fromIp, srcPort, destIp, destPort, string(payload))
		}
	}

}

func checkUsername(payload []byte) (string, bool) {
	field := ""
	result := false
	for _, u := range userList {
		payload = []byte(strings.ToLower(string(payload)))
		if bytes.Contains(payload, []byte(strings.ToLower(u))) {
			field = u
			result = true
			break
		}
	}

	return field, result
}

func checkPassword(payload []byte) (string, bool) {
	field := ""
	result := false
	for _, p := range passList {
		payload = []byte(strings.ToLower(string(payload)))
		if bytes.Contains(payload, []byte(strings.ToLower(p))) {
			field = p
			result = true
			break
		}
	}

	return field, result
}
