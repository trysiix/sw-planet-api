package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Planet -  structure for planets data
type Planet struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	Name                string             `bson:"name,omitempty"`
	Weather             string             `bson:"weather,omitempty"`
	Terrain             string             `bson:"terrain,omitempty"`
	NumberOfAppearances int                `bson:"numberofappearances,omitempty"`
}

// Swapi - structure to handle swapi requested data
type Swapi struct {
	Count    int       `json:"count,omitempty"`
	Next     int       `json:"next,omitempty"`
	Previous int       `json:"previous,omitempty"`
	Results  []Results `json:"results,omitempty"`
}

//Results - structure to handle the results data
type Results struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	ResidentURLs   []string `json:"residents"`
	FilmURLs       []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}
