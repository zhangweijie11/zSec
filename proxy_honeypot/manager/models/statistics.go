package models

import (
	"context"
	"github.com/zhangweijie11/zSec/proxy_honeypot/manager/vars"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// TongjiPasswordBySite 统计每个网站的密码数量
func TongjiPasswordBySite(page int) (passwords []bson.M, pages int, total int, err error) {
	coll := Database.Collection("password")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 创建聚合管道
	pipeline := mongo.Pipeline{
		{{"$group", bson.D{{"_id", "$host"}, {"count", bson.D{{"$sum", 1}}}}}},
		{{"$sort", bson.D{{"count", -1}}}},
	}

	// 执行聚合查询
	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, 0, err
	}
	defer cursor.Close(ctx)

	// 读取所有数据到内存中，如果数据量大则不推荐这样做，会有性能问题
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, 0, 0, err
	}

	total = len(results)
	pages = calculatePages(total, vars.PageSize)

	start, end := calculatePageRange(page, pages, vars.PageSize, total)
	passwords = results[start:end]

	return passwords, pages, total, nil
}

// TongjiUrls 统计每个host的访问次数
func TongjiUrls(page int) (urls []bson.M, pages int, total int, err error) {
	coll := Database.Collection("proxy_honeypot")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 创建聚合管道
	pipeline := mongo.Pipeline{
		{{"$group", bson.D{{"_id", "$host"}, {"count", bson.D{{"$sum", 1}}}}}},
		{{"$sort", bson.D{{"count", -1}}}},
	}

	opts := options.Aggregate().SetAllowDiskUse(true) // 允许使用磁盘缓存

	// 执行聚合查询
	cursor, err := coll.Aggregate(ctx, pipeline, opts)
	if err != nil {
		return nil, 0, 0, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, 0, 0, err
	}

	total = len(results)
	pages = calculatePages(total, vars.PageSize)

	start, end := calculatePageRange(page, pages, vars.PageSize, total)
	urls = results[start:end]

	return urls, pages, total, nil
}

// calculatePages 计算总页数
func calculatePages(total, pageSize int) int {
	if total%pageSize == 0 {
		return total / pageSize
	}
	return total/pageSize + 1
}

// calculatePageRange 计算页面范围
func calculatePageRange(page, pages, pageSize, total int) (start, end int) {
	if page < 1 {
		page = 1
	} else if page > pages {
		page = pages
	}
	start = (page - 1) * pageSize
	end = start + pageSize
	if end > total {
		end = total
	}
	return start, end
}
