package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Dns struct {
	DnsType string `json:"dns_type"`
	DnsName string `json:"dns_name"`
	SrcIp   string `json:"src_ip"`
	DstIp   string `json:"dst_ip"`
}

type EvilDns struct {
	Id       int64
	Time     time.Time `bson:"time"`
	SensorIp string    `bson:"sensor_ip"`
	IsEvil   bool      `bson:"is_evil"`
	Data     Dns       `bson:"data"`
}

func NewEvilDns(sensorIp string, isEvil bool, dns Dns) (evilDns *EvilDns) {
	now := time.Now()
	return &EvilDns{SensorIp: sensorIp, Time: now, IsEvil: isEvil, Data: dns}
}

func (d *EvilDns) Insert() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database(DbConfig.DbName).Collection("dns")
	_, err := collection.InsertOne(ctx, d)
	return err
}

func ListEvilDns() ([]EvilDns, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database(DbConfig.DbName).Collection("dns")

	opts := options.Find().SetSort(bson.D{{"_id", -1}}).SetLimit(500)
	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []EvilDns
	for cursor.Next(ctx) {
		var dns EvilDns
		if err := cursor.Decode(&dns); err != nil {
			return nil, err
		}
		results = append(results, dns)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
