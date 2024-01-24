package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// DashPassword retrieves top 20 passwords by site
func DashPassword() (passwords []bson.M, err error) {
	coll := Database.Collection("password")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipe := []bson.M{
		{"$group": bson.M{"_id": "$site", "count": bson.M{"$sum": 1}}},
		{"$sort": bson.M{"count": -1}},
		{"$limit": 20},
	}

	cursor, err := coll.Aggregate(ctx, pipe)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &passwords); err != nil {
		return nil, err
	}
	return passwords, nil
}

// DashUrls retrieves the latest 20 URLs
func DashUrls() (urls []bson.M, err error) {
	coll := Database.Collection("urls")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetLimit(20)
	cursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &urls); err != nil {
		return nil, err
	}
	return urls, nil
}

// DashIps retrieves the latest 20 evil ips
func DashIps() (evilIps []bson.M, err error) {
	coll := Database.Collection("evil_ips")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetLimit(20)
	cursor, err := coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &evilIps); err != nil {
		return nil, err
	}
	return evilIps, nil
}

// DashTotal retrieves the total number of records and passwords
func DashTotal() (totalRecord int64, totalPassword int64, err error) {
	coll := Database.Collection("proxy_honeypot")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	totalRecord, err = coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, 0, err
	}

	collPassword := Database.Collection("password")
	totalPassword, err = collPassword.CountDocuments(ctx, bson.M{})
	if err != nil {
		return totalRecord, 0, err
	}
	return totalRecord, totalPassword, nil
}
