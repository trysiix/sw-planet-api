package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"

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

// Add creates the 'add' route to register the planet info, executes on call
func Add(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var planet models.Planet
	_ = json.NewDecoder(r.Body).Decode(&planet)

	fmt.Println(planet, r.Body)

	addPlanet(planet)

	json.NewEncoder(writer).Encode(planet)
}

// Insert the planet data into mongo db
func addPlanet(planet models.Planet) {
	insertResult, err := collection.InsertOne(context.Background(), planet)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Planet Added", insertResult.InsertedID)
}
