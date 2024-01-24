package models

import (
	"context"
	"fmt"
	"github.com/zhangweijie11/zSec/proxy_honeypot/manager/logger"
	"github.com/zhangweijie11/zSec/proxy_honeypot/manager/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
	Host     string
	Port     int
	USERNAME string
	PASSWORD string
	DataName string

	collAdmin *mongo.Collection
)

func init() {
	cfg := settings.Cfg
	sec := cfg.Section("MONGODB")
	Host = sec.Key("HOST").MustString("127.0.0.1")
	Port = sec.Key("PORT").MustInt(27017)
	USERNAME = sec.Key("USER").MustString("xproxy")
	PASSWORD = sec.Key("PASS").MustString("passw0rd")
	DataName = sec.Key("DATA").MustString("xproxy")

	if err := NewMongodbClient(); err != nil {
		logger.Logger.Panicf("CONNECT MONGODB, err: %v", err)
	}

	collAdmin = Database.Collection("users")
	userCount, err := collAdmin.EstimatedDocumentCount(context.Background())
	if err != nil {
		logger.Logger.Errorf("Error getting user count: %v", err)
	}
	if userCount == 0 {
		if err := NewUser("xproxy", "x@xsec.io"); err != nil {
			logger.Logger.Errorf("Error creating new user: %v", err)
		}
	}
}

// NewMongodbClient returns a new mongodb client
func NewMongodbClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", USERNAME, PASSWORD, Host, Port, DataName)
	clientOptions := options.Client().ApplyURI(uri)

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = Client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	Database = Client.Database(DataName)
	return nil
}
