package sniffer

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/toolkits/slice"
	"github.com/zhangweijie11/zSec/honeypot/agent/logger"
	"github.com/zhangweijie11/zSec/honeypot/agent/models"
	"strings"
	"sync"
	"time"
)

var (
	mux sync.Mutex
)

func ProcessPacket(packets chan gopacket.Packet) {
	for packet := range packets {
		processPacket(packet)
	}
}

func processPacket(packet gopacket.Packet) {
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		ip, ok := ipLayer.(*layers.IPv4)
		if ok {
			switch ip.Protocol {
			case layers.IPProtocolTCP:
				tcpLayer := packet.Layer(layers.LayerTypeTCP)
				if tcpLayer != nil {
					tcp, _ := tcpLayer.(*layers.TCP)

					srcPort := SplitPortService(tcp.SrcPort.String())
					dstPort := SplitPortService(tcp.DstPort.String())
					isHttp := false

					applicationLayer := packet.ApplicationLayer()
					if applicationLayer != nil {
						// Search for a string inside the payload
						if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
							isHttp = true
						}
					}

					connInfo := models.NewConnectionInfo("tcp", ip.SrcIP.String(), srcPort, ip.DstIP.String(), dstPort, isHttp)

					go func(info *models.ConnectionInfo) {
						if !IsInWhite(info) &&
							!CheckSelfPacker(info) &&
							(tcp.SYN && !tcp.ACK) {
							err := SendPacker(info)
							logger.Log.Debugf("[TCP] %v:%v -> %v:%v, err: %v", ip.SrcIP, tcp.SrcPort.String(),
								ip.DstIP, tcp.DstPort.String(), err)
						}
					}(connInfo)
				}
			}
		}
	}
}

func SendPacker(connInfo *models.ConnectionInfo) (err error) {
	packetInfo := models.NewPacketInfo(connInfo, time.Now())
	jsonPacket, err := packetInfo.String()
	if err != nil {
		return err
	}

	go logger.LogReport.WithField("api", "/api/packet/").Info(jsonPacket)

	return err
}

func CheckSelfPacker(p *models.ConnectionInfo) (ret bool) {
	if slice.ContainsString(SensorIps, p.SrcIp) || p.DstIp == ApiIp || p.SrcIp == ApiIp {
		ret = true
	}
	return ret
}

func SplitPortService(portService string) (port string) {
	t := strings.Split(portService, "(")
	if len(t) > 0 {
		port = t[0]
	}
	return port
}
