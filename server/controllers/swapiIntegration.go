package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../models"
)

const sniffLen = 512

// GetPlanetsData test
func GetPlanetsData(Planet string) []byte {

	url := "http://swapi.dev/api/planets/?search=" + Planet

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body
}

// GetNumOfAppearances this func gets the number of appearances in films
func GetNumOfAppearances(Planet string) int {

	var data models.Swapi

	jsonStr := string(GetPlanetsData(Planet))
	json.Unmarshal([]byte(jsonStr), &data)
	numOfAppearances := len(data.Results[0].FilmURLs)

	return numOfAppearances
}