package database

import (
	"context"
	"fmt"
	"log"

	"github.com/andreposman/zipcode-producer/pkg/settings"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetClient creates a client
func GetClient() *options.ClientOptions {
	connectionString := settings.CreateConnectionString()
	clientOptions := options.Client().ApplyURI(connectionString)

	return clientOptions
}

//Connect to mongo
func Connect() {
	clientOptions := GetClient()
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	//check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDb")
}
