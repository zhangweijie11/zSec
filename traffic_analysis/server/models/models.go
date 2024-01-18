package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

var (
	// Assume the Client is already connected and available
	collection *mongo.Collection
)

type ConnectionInfo struct {
	Protocol string `json:"protocol"`
	SrcIp    string `json:"src_ip"`
	SrcPort  string `json:"src_port"`
	DstIp    string `json:"dst_ip"`
	DstPort  string `json:"dst_port"`
}

// bad ip or dns source info
type Source struct {
	Desc   string `json:"desc"`
	Source string `json:"source"`
}

// evil ips
type EvilIps struct {
	Ips []string `json:"ips"`
	Src Source   `json:"src"`
}

type IpList struct {
	Id   int64
	Ip   string   `json:"ip"`
	Info []Source `json:"info"`
}

type IplistApi struct {
	Evil bool   `json:"evil"`
	Data IpList `json:"data"`
}

type EvilConnectInfo struct {
	Id       int64
	Time     time.Time `bson:"time"`
	SensorIp string    `bson:"sensor_ip"`
	Protocol string    `bson:"protocol"`
	SrcIp    string    `bson:"src_ip"`
	SrcPort  string    `bson:"src_port"`
	DstIp    string    `bson:"dst_ip" `
	DstPort  string    `bson:"dst_port" `
	IsEvil   bool      `bson:"is_evil" `
	Data     []Source  `bson:"data"`
}

func init() {
	collection = Client.Database(DbConfig.DbName).Collection("connection_info")
}

func NewEvilConnectionInfo(sensorIp string, info ConnectionInfo, evilData IplistApi) (evilInfo *EvilConnectInfo) {
	now := time.Now()
	return &EvilConnectInfo{SensorIp: sensorIp, Time: now, Protocol: info.Protocol, SrcIp: info.SrcIp,
		SrcPort: info.SrcPort, DstIp: info.DstIp, DstPort: info.DstPort, IsEvil: evilData.Evil, Data: evilData.Data.Info}
}

func (i *EvilConnectInfo) Insert() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exists, err := i.Exist(ctx)
	if err != nil {
		return err
	}
	if !exists {
		_, err = collection.InsertOne(ctx, i)
	}
	return err
}

func (i *EvilConnectInfo) Exist(ctx context.Context) (bool, error) {
	srcIp := strings.Split(i.SrcIp, ":")[0]
	filter := bson.M{"src_ip": srcIp, "dst_ip": i.DstIp}

	var result EvilConnectInfo
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	return err == nil, err
}

func ListEvilInfo(ctx context.Context) ([]EvilConnectInfo, error) {
	var evilInfos []EvilConnectInfo

	opts := options.Find().SetSort(bson.D{{"_id", -1}}).SetLimit(500)
	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var evilInfo EvilConnectInfo
		if err = cursor.Decode(&evilInfo); err != nil {
			return nil, err
		}
		evilInfos = append(evilInfos, evilInfo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return evilInfos, nil
}
