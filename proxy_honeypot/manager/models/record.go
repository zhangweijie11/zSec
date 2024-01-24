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

type HttpRecord struct {
	Id            primitive.ObjectID `bson:"_id"`
	Session       int64              `json:"session"`
	Method        string             `json:"method"`
	RemoteAddr    string             `json:"remote_addr" bson:"remote"`
	StatusCode    int                `json:"status"`
	ContentLength int64              `json:"content_length"`
	Host          string             `json:"host"`
	Port          string             `json:"port"`
	Url           string             `json:"url"`
	Scheme        string             `json:"scheme"`
	Path          string             `json:"path"`
	ReqHeader     http.Header        `json:"req_header"`
	RespHeader    http.Header        `json:"resp_header"`
	RequestParam  url.Values         `json:"request_param" bson:"requestparameters"`
	RequestBody   []byte             `json:"request_body"`
	ResponseBody  []byte             `json:"response_body"`
	VisitTime     time.Time          `json:"visit_time"`
}

// ListRecordByPage returns a paginated slice of HttpRecord, the total number of pages and the total number of records
func ListRecordByPage(page int) (records []HttpRecord, pages int, total int64, err error) {
	coll := Database.Collection("record")
	ctx := context.Background()

	// Calculate the total number of records
	total, err = coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, 0, err
	}

	// Calculate the total number of pages
	recordPerPage := int64(vars.PageSize)
	pages = int(total) / vars.PageSize
	if total%recordPerPage > 0 {
		pages++
	}

	// Check if the requested page is within the available range
	if page > pages {
		page = pages
	}
	if page < 1 {
		page = 1
	}

	// Set the options for pagination
	option := options.Find().SetSkip(recordPerPage * int64(page-1)).SetLimit(recordPerPage)

	// Find the records for the requested page
	cursor, err := coll.Find(ctx, bson.M{}, option)
	if err != nil {
		return nil, 0, 0, err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document
	for cursor.Next(ctx) {
		var record HttpRecord
		if err := cursor.Decode(&record); err != nil {
			return nil, 0, 0, err
		}
		records = append(records, record)
	}

	return records, pages, total, err
}

// ListRecordBySite returns a slice of HttpRecord by given host and page number
func ListRecordBySite(site string, page int) (records []HttpRecord, pages int, total int64, err error) {
	coll := Database.Collection("record")
	filter := bson.M{"host": site}
	total, err = coll.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, 0, err
	}

	pages = int(total) / vars.PageSize
	if total%int64(vars.PageSize) > 0 {
		pages++
	}

	if page > pages {
		page = pages
	}
	if page < 1 {
		page = 1
	}

	options := options.Find().SetSkip(int64(vars.PageSize * (page - 1))).SetLimit(int64(vars.PageSize))
	cursor, err := coll.Find(context.Background(), filter, options)
	if err != nil {
		return nil, 0, 0, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var record HttpRecord
		if err := cursor.Decode(&record); err != nil {
			return nil, 0, 0, err
		}
		records = append(records, record)
	}

	return records, pages, int64(int(total)), err
}

// RecordDetail returns a single HttpRecord by its ID
func RecordDetail(id string) (HttpRecord, error) {
	var record HttpRecord
	coll := Database.Collection("record")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return record, err
	}

	err = coll.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&record)
	return record, err
}
