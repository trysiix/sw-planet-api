package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// Create , creates the route to register the planet info, executes on call
func Create(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var planet models.Planet
	_ = json.NewDecoder(r.Body).Decode(&planet)

	fmt.Println(planet, r.Body)

	create(planet)

	json.NewEncoder(writer).Encode(planet)
}

// Insert the planet data into mongo db
func create(planet models.Planet) {
	insertResult, err := collection.InsertOne(context.Background(), planet)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Planet Added", insertResult.InsertedID)
}

// IndexAll builds the index route and execute on call
func IndexAll(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "GET")

	payload := indexAll()

	json.NewEncoder(writer).Encode(payload)
}

// indexAll - gets all the planets registered in mongo db
func indexAll() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// DeleteByID , this is the delete function route
func DeleteByID(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "DELETE")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteByID(params["id"])
	json.NewEncoder(writer).Encode(params["id"])

}

// delete planet by ID
func deleteByID(ID string) {
	fmt.Println(ID)
	id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Planet", d.DeletedCount)
}
