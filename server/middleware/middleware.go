package middleware

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
const connectionString = "mongodb+srv://admin:QUTjF2P68cZADUx0@c-zero.aju4h.gcp.mongodb.net/<dbname>?retryWrites=true&w=majority"

// Database Name
const dbName = "Skyriver"

// Collection name
const collName = "planet"

// collection object/instance
var collection *mongo.Collection

// init - create connection with mongo db
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Creating collection instance
	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")

}
