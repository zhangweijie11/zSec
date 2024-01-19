package models

import (
	"encoding/json"
	"time"
)

type (
	ConnectionInfo struct {
		Protocol string `json:"protocol"`
		SrcIp    string `json:"src_ip"`
		SrcPort  string `json:"src_port"`
		DstIp    string `json:"dst_ip"`
		DstPort  string `json:"dst_port"`
		IsHttp   bool   `json:"is_http"`
	}

	PacketInfo struct {
		ConnInfo *ConnectionInfo `json:"conn_info"`
		Time     time.Time       `json:"time"`
	}
)

func NewConnectionInfo(proto string, srcIp string, srcPort string, dstIp string, dstPort string, isHttp bool) (connInfo *ConnectionInfo) {
	return &ConnectionInfo{Protocol: proto, SrcIp: srcIp, SrcPort: srcPort, DstIp: dstIp, DstPort: dstPort, IsHttp: isHttp}
}

func (c *ConnectionInfo) String() (string, error) {
	js, err := json.Marshal(c)
	return string(js), err
}

func NewPacketInfo(info *ConnectionInfo, now time.Time) (ret *PacketInfo) {
	ret = &PacketInfo{ConnInfo: info, Time: now}
	return ret
}

func (p *PacketInfo) String() (string, error) {
	js, err := json.Marshal(p)
	return string(js), err
}
