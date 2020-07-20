package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"../controllers"
	"../models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Declaring control consts
const (
	connectionString = "mongodb+srv://admin:fftddo90Sk5VPuRb@cluster0.aju4h.gcp.mongodb.net/Skyriver?retryWrites=true&w=majority"
	dbName           = "Skyriver"
	collName         = "planets"
	AllowedOrigin    = "Access-Control-Allow-Origin"
	AllowedMethods   = "Access-Control-Allow-Methods"
	AllowedHeaders   = "Access-Control-Allow-Headers"
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
	w.Header().Set(AllowedOrigin, "*")
	w.Header().Set(AllowedMethods, "POST")
	w.Header().Set(AllowedHeaders, "Content-Type")

	var data models.Planet
	_ = json.NewDecoder(r.Body).Decode(&data)
	params := &data
	if len(params.Name) > 0 && len(params.Terrain) > 0 && len(params.Weather) > 0 {
		fmt.Println(params)
		data.NumberOfAppearances = controllers.GetNumOfAppearances(params.Name)

		data.Name = strings.ToLower(data.Name)
		data.Weather = strings.ToLower(data.Weather)
		data.Terrain = strings.ToLower(data.Terrain)

		_, err := collection.InsertOne(context.Background(), data)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(data)
	} else {
		alert := []byte("Fields incorrectly filled")
		w.Write(alert)
	}
}

// Index - Index the planets by request, with filters or without
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set(AllowedOrigin, "*")
	w.Header().Set(AllowedMethods, "GET")

	keys := r.URL.Query()
	payload := []primitive.M{}
	for fName, fValue := range keys {
		if len(fName) > 0 && len(fValue[0]) > 0 {
			payload = getDocuments(fName, fValue[0])

		} else {
			alert := []byte("Invalid Filter, Ex: /api/planet?name=Naboo")
			w.Write(alert)
			return
		}
	}

	if len(payload) <= 0 {
		payload = getDocuments("", "")
	}

	json.NewEncoder(w).Encode(payload)
}

// getDocument - gets all the planets registered
func getDocuments(fName string, fValue string) []primitive.M {
	filter := bson.M{}

	fName = strings.ToLower(fName)
	fValue = strings.ToLower(fValue)

	if len(fName) > 0 && len(fValue) > 0 {
		filter = bson.M{fName: fValue}
	}

	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// IndexByID - Indexes a registered planet by the provided id
func IndexByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set(AllowedOrigin, "*")
	w.Header().Set(AllowedMethods, "GET")

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
	w.Header().Set(AllowedOrigin, "*")
	w.Header().Set(AllowedMethods, "DELETE")
	w.Header().Set(AllowedHeaders, "Content-Type")

	params := mux.Vars(r)
	docID := params["id"]
	id, _ := primitive.ObjectIDFromHex(docID)
	filter := bson.M{"_id": id}
	deleted, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if deleted.DeletedCount > 0 {
		json.NewEncoder(w).Encode("Deleted Document: " + docID)

	} else {
		alert := []byte("We couldn't find any document with this id: " + docID)
		w.Write(alert)
	}
}
