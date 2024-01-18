package models

import (
	"context"
	"fmt"
	"github.com/zhangweijie11/zSec/traffic_analysis/sensor/settings"
	"github.com/zhangweijie11/zSec/traffic_analysis/server/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	DbConfig DbCONF
	Client   *mongo.Client
)

type DbCONF struct {
	DbType string
	DbHost string
	DbPort int64
	DbUser string
	DbPass string
	DbName string
}

func init() {
	cfg := settings.Cfg
	sec := cfg.Section("database")
	DbConfig.DbType = sec.Key("DB_TYPE").MustString("mongodb")
	DbConfig.DbHost = sec.Key("DB_HOST").MustString("127.0.0.1")
	DbConfig.DbPort = sec.Key("DB_PORT").MustInt64(27017)
	DbConfig.DbUser = sec.Key("DB_USER").MustString("user")
	DbConfig.DbPass = sec.Key("DB_PASS").MustString("password")
	DbConfig.DbName = sec.Key("DB_NAME").MustString("proxy_honeypot")
	if err := NewDbEngine(); err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
}

func NewDbEngine() error {
	var err error
	ctx := context.TODO()

	// 构建连接串
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		DbConfig.DbUser,
		DbConfig.DbPass,
		DbConfig.DbHost,
		DbConfig.DbPort,
		DbConfig.DbName,
	)

	// 设置客户端选项
	clientOptions := options.Client().ApplyURI(uri)

	// 连接到MongoDB
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		util.Log.Panicf("Connect Database failed, err: %v", err)
		return err
	}

	// 检查连接
	err = Client.Ping(ctx, nil)
	if err != nil {
		util.Log.Panicf("Failed to connect to database, err: %v", err)
		return err
	}

	util.Log.Infof("DB Type: %v, Connect err status: %v", DbConfig.DbType, err)

	return nil
}
