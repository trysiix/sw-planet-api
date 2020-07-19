package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../controllers"
	"../models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
const connectionString = "mongodb+srv://admin:fftddo90Sk5VPuRb@cluster0.aju4h.gcp.mongodb.net/Skyriver?retryWrites=true&w=majority"

// Database Name
const dbName = "Skyriver"

// Collection name
const collName = "planets"

var collection *mongo.Collection

// init - creates the connection with mongo db
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

	var data models.Planet
	_ = json.NewDecoder(r.Body).Decode(&data)

	params := &data

	data.NumberOfAppearances = controllers.GetNumOfAppearances(params.Name)

	fmt.Println(data)

	insertResult, err := collection.InsertOne(context.Background(), data)
	fmt.Println("Planet Save", insertResult)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(writer).Encode(data)
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

// IndexByID builds the index route and execute on call
func IndexByID(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "GET")

	params := mux.Vars(r)

	fmt.Println(params)
	payload := indexByID(params["id"])
	json.NewEncoder(writer).Encode(payload)
}

// indexByID - gets all the planets registered in mongo db
func indexByID(item string) interface{} {

	docID, err := primitive.ObjectIDFromHex(item)

	result := models.Planet{}

	filter := bson.M{"_id": docID}

	err = collection.FindOne(context.Background(), filter).Decode(&result)

	fmt.Println(result)

	if err != nil {
		log.Fatal(err)
	}

	return result

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
