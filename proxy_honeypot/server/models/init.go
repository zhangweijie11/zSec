package models

import (
	"context"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/zhangweijie11/zSec/proxy_honeypot/server/log"
	"github.com/zhangweijie11/zSec/proxy_honeypot/server/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

var (
	DbConfig    DbCONF
	Engine      *xorm.Engine
	MongoClient *mongo.Client
)

type DbCONF struct {
	DbType string
	DbHost string
	DbPort int64 // MongoDB端口通常作为字符串保存
	DbUser string
	DbPass string
	DbName string
}

func init() {
	cfg := settings.Cfg
	sec := cfg.Section("database")
	DbConfig.DbType = sec.Key("DB_TYPE").MustString("mysql")
	DbConfig.DbHost = sec.Key("DB_HOST").MustString("127.0.0.1")
	DbConfig.DbPort = sec.Key("DB_PORT").MustInt64(3306)
	DbConfig.DbUser = sec.Key("DB_USER").MustString("x-proxy")
	DbConfig.DbPass = sec.Key("DB_PASS").MustString("x@xsec.io")
	DbConfig.DbName = sec.Key("DB_NAME").MustString("x-proxy")

	_ = NewDbEngine()

}

func NewDbEngine() (err error) {
	switch DbConfig.DbType {
	case "mysql":
		dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
			DbConfig.DbUser, DbConfig.DbPass, DbConfig.DbHost, strconv.Itoa(int(DbConfig.DbPort)), DbConfig.DbName)
		Engine, err = xorm.NewEngine("mysql", dataSourceName)
		if err == nil {
			err = Engine.Ping()
			if err == nil {
				_ = Engine.Sync2(new(Record))
			}
		}

	case "mongodb":
		var err error
		MongoClient, err = NewMongoClient()
		if err != nil {
			log.Logger.Infof("Failed to connect to MongoDB: ", err)
		}
		// 确认连接
		err = MongoClient.Ping(context.Background(), nil)
		if err != nil {
			log.Logger.Infof("Failed to ping MongoDB: ", err)
		}
		log.Logger.Infof("Connected to MongoDB")
	}

	return err
}

func GetMongoCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database(DbConfig.DbName).Collection(collectionName)
}

func NewMongoClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
			DbConfig.DbUser, DbConfig.DbPass, DbConfig.DbHost, strconv.Itoa(int(DbConfig.DbPort)), DbConfig.DbName)).
		SetAuth(options.Credential{
			AuthSource: DbConfig.DbName, // "AuthSource" usually defaults to the database name
			Username:   DbConfig.DbUser,
			Password:   DbConfig.DbPass,
		})

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %v", err)
	}

	return client, nil
}
