package mongo_ops

import (
	"fmt"
	"log"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"backend.com/example/go-backend/json_structs"
	
)

var Client * mongo.Client

func ConnectMongo() bool {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		fmt.Println(time.Now())	
		return false
	}
	return true
}


func DisconnectMongo()  {
	if  err := Client.Disconnect(context.TODO()); err != nil{
		log.Fatal(err)
	}
}


func SignupMethod(signup json_structs.SignupDetails ) bool {
	// we will be defining this for every method that we create
	// dbname := "stanza_copy"
	
	client := Client.Database("stanza_copy").Collection("users")



	insDoc := bson.D{
		{Key :"mail" , Value: signup.Email},
		{Key :"username" , Value: signup.Username},
		{Key :"password" , Value: signup.Password},
	}
	// this is creating the json for find query
	signupJson := bson.D{
		// for array we use bson.A as or query needs array to check
		{Key: "$or", Value: bson.A{
			bson.D{{Key :"mail" , Value: signup.Email}},
			bson.D{{Key :"username" , Value : signup.Username}},
		}},
	}
	var result interface{}
	err := client.FindOne(context.TODO(),signupJson).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return true
	}
	if result == nil{
		res, err := client.InsertOne(context.TODO(), insDoc)
		if err != nil{
			fmt.Println(err)
		}
		_ = res
		return false
	}
	
	return false
}