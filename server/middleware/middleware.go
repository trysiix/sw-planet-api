package middleware

import (
	"context"
	"encoding/json"
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

//Declaring mongo db settings to connect
const (
	connectionString = "mongodb+srv://admin:fftddo90Sk5VPuRb@cluster0.aju4h.gcp.mongodb.net/Skyriver?retryWrites=true&w=majority"
	dbName           = "Skyriver"
	collName         = "planets"
)

var collection *mongo.Collection

// init - Connects to mongo db and builds instance of collection
func init() {

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(collName)

}

// Create - Register the planet data
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var data models.Planet
	_ = json.NewDecoder(r.Body).Decode(&data)

	params := &data
	data.NumberOfAppearances = controllers.GetNumOfAppearances(params.Name)

	_, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(data)
}

// IndexAll - Index all the planets registerd
func IndexAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

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

	json.NewEncoder(w).Encode(results)
}

// IndexByID - Indexes a registered planet by the provided id
func IndexByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	params := mux.Vars(r)

	docID, err := primitive.ObjectIDFromHex(params["id"])

	result := models.Planet{}

	filter := bson.M{"_id": docID}

	err = collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

// DeleteByID - Deletes a planet document by the provided id
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(params["id"])
}
