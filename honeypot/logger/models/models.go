package models

import (
	"context"
	"fmt"
	"github.com/zhangweijie11/zSec/honeypot/logger/settings"
	"github.com/zhangweijie11/zSec/honeypot/logger/vars"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	Client            *mongo.Client
	CollectionPacket  *mongo.Collection
	CollectionService *mongo.Collection
)

func ConnectMongoDb() (*mongo.Client, error) {
	Cfg := settings.Cfg
	sec := Cfg.Section("database")
	host := sec.Key("HOST").MustString("127.0.0.1")
	port := sec.Key("PORT").MustInt(27017)
	user := sec.Key("USER").MustString("honeypot")
	password := sec.Key("PASSWORD").MustString("zSec")
	database := sec.Key("DATABASE").MustString("zSec-honeypot")

	vars.MongodbName = database
	vars.CollPacket = sec.Key("COLL_PACKET").MustString("packet_info")
	vars.CollService = sec.Key("COLL_SERVICE").MustString("service")

	mongodbUrl := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v",
		user,
		password,
		host,
		port,
		database,
	)

	clientOptions := options.Client().ApplyURI(mongodbUrl)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func init() {
	err := checkConnect()
	if err != nil {
		log.Panic(err)
	}
}

func checkConnect() error {
	if Client == nil {
		client, err := ConnectMongoDb()
		if err != nil {
			return err
		}
		Client = client
	}

	database := Client.Database(vars.MongodbName)
	CollectionService = database.Collection(vars.CollService)
	CollectionPacket = database.Collection(vars.CollPacket)

	return nil
}
