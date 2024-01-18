package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/url"
	"time"
)

type HttpReq struct {
	Host          string
	Ip            string
	Client        string
	Port          string
	URL           *url.URL
	Header        http.Header
	RequestURI    string
	Method        string
	ReqParameters url.Values
}

type EvilHttpReq struct {
	Id       int64
	Time     time.Time `bson:"time"`
	SensorIp string    `bson:"sensor_ip"`
	IsEvil   bool      `bson:"is_evil"`
	Data     HttpReq   `bson:"data"`
}

func NewEvilHttpReq(sensorIp string, isEvil bool, req HttpReq) (evilHttpReq *EvilHttpReq) {
	now := time.Now()
	return &EvilHttpReq{SensorIp: sensorIp, Time: now, IsEvil: isEvil, Data: req}
}

// Insert方法插入一个新的EvilHttpReq到http_req集合
func (e *EvilHttpReq) Insert() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 假设Client是一个已经连接的mongo客户端实例
	collection := Client.Database(DbConfig.DbName).Collection("http_req")
	_, err := collection.InsertOne(ctx, e)
	return err
}

// ListEvilHttpReq方法列出http_req集合中的最多500条记录
func ListEvilHttpReq() ([]EvilHttpReq, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := Client.Database(DbConfig.DbName).Collection("http_req")

	// 创建一个FindOptions，并设置结果的排序和数量限制
	opts := options.Find().SetSort(bson.D{{"_id", -1}}).SetLimit(500)

	// 执行查询操作
	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var evilHttpReqs []EvilHttpReq
	// 遍历查询结果
	for cursor.Next(ctx) {
		var httpReq EvilHttpReq
		if err = cursor.Decode(&httpReq); err != nil {
			return nil, err
		}
		evilHttpReqs = append(evilHttpReqs, httpReq)
	}

	// 检查遍历过程中是否有错误发生
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return evilHttpReqs, nil
}
