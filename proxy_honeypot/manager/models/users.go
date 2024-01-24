package models

import (
	"context"
	"github.com/zhangweijie11/zSec/proxy_honeypot/manager/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	UserName string             `bson:"username"`
	Password string             `bson:"password"`
}

func ListUser() (users []User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collAdmin.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id string) (user User, err error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collAdmin.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	return user, err
}

func NewUser(username, password string) (err error) {
	encryptPass := util.EncryptPass(password)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collAdmin.InsertOne(ctx, User{Id: primitive.NewObjectID(), UserName: username, Password: encryptPass})
	return err
}

func UpdateUser(id string, username, password string) (err error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"username": username,
			"password": util.EncryptPass(password),
		},
	}
	_, err = collAdmin.UpdateByID(ctx, objID, update)
	return err
}

func DelUser(id string) (err error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collAdmin.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

func Auth(username, password string) (result bool, err error) {
	encryptPass := util.EncryptPass(password)
	userAuth := User{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collAdmin.FindOne(ctx, bson.M{"username": username, "password": encryptPass}).Decode(&userAuth)
	if err == nil && userAuth.UserName == username && userAuth.Password == encryptPass {
		result = true
	}
	return result, err
}
