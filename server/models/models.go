package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Planet - Creating data structure
type Planet struct {
	ID          primitive.ObjectID
	Name        string
	Weather     string
	TerrainType string
}
