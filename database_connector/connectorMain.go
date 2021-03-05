package database_connector

import(
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"context"
)

var dbCollection *mongo.Collection
var ctx = context.TODO()

func Connect(mongoUri string) *mongo.Client {

	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {

		log.Printf("Successful connection established to the database.")

	}
	return client

}
