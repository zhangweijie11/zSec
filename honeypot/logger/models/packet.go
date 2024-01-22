package models

import (
	"context"
	"time"
)

type (
	ConnectionInfo struct {
		Protocol string `json:"protocol,omitempty"`
		SrcIp    string `json:"src_ip,omitempty"`
		SrcPort  string `json:"src_port,omitempty"`
		DstIp    string `json:"dst_ip,omitempty"`
		DstPort  string `json:"dst_port,omitempty"`
		IsHttp   bool   `json:"is_http,omitempty"`
	}

	PacketInfo struct {
		ConnInfo *ConnectionInfo `json:"conn_info"`
		Time     time.Time       `json:"time"`
	}
)

func (p *PacketInfo) Insert() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := CollectionPacket.InsertOne(ctx, p)
	return err
}
