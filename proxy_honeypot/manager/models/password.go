package models

import (
	"context"
	"github.com/zhangweijie11/zSec/proxy_honeypot/manager/vars"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"net/url"
	"time"
)

type Password struct {
	Id                primitive.ObjectID `bson:"_id"`
	ResponseBody      string             `bson:"response_body"`
	RequestBody       string             `bson:"request_body"`
	DateStart         time.Time          `bson:"date_start"`
	URL               string             `bson:"url"`
	RequestParameters url.Values         `bson:"request_parameters"`
	FromIp            string             `bson:"from_ip"`
	Site              string             `bson:"site"`
	ResponseHeader    http.Header        `bson:"response_header"`
	RequestHeader     http.Header        `bson:"request_header"`
	Data              map[string]string  `bson:"data"`
}

func ListPasswordByPage(page int) (passwords []Password, pages int, total int64, err error) {
	coll := Database.Collection("password")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	total, err = coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, 0, err
	}

	pages = int(total) / vars.PageSize
	if int(total)%vars.PageSize != 0 {
		pages++
	}

	if page > pages {
		page = pages
	}
	if page < 1 {
		page = 1
	}

	skip := (page - 1) * vars.PageSize
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(vars.PageSize))
	cursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, total, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &passwords); err != nil {
		return nil, 0, total, err
	}

	return passwords, pages, total, nil
}

func ListPasswordBySite(site string, page int) (passwords []Password, pages int, total int64, err error) {
	filter := bson.M{"site": site}
	coll := Database.Collection("password")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	total, err = coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, 0, err
	}

	pages = int(total) / vars.PageSize
	if int(total)%vars.PageSize != 0 {
		pages++
	}

	if page > pages {
		page = pages
	}
	if page < 1 {
		page = 1
	}

	skip := (page - 1) * vars.PageSize
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(vars.PageSize))
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, total, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &passwords); err != nil {
		return nil, 0, total, err
	}

	return passwords, pages, total, nil
}

func PasswordDetail(id string) (Password, error) {
	var password Password
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return password, err
	}

	filter := bson.M{"_id": objID}
	coll := Database.Collection("password")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = coll.FindOne(ctx, filter).Decode(&password)
	return password, err
}
