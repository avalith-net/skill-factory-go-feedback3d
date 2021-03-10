package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var dbURI = os.Getenv("DB_URI")

/*MongoCN is the object of connecting to the database*/
var MongoCN = ConnectionDB()
var clientOptions = options.Client().ApplyURI(os.Getenv("DB_URI"))

/*ConnectionDB is the feature that allows me to connect to the database
  Returns a connection to the BD of type Mongo Client*/
func ConnectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Successful connection to BD")
	return client
}

/*CheckConnection it's ping to BD*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
