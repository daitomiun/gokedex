package service

import (
	"encoding/json"
	. "github.com/daitomiun/gokedex/models"
	"log"
	"net/http"
)

func GetMapLocations(url string) (Config, []string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	var mapLocations []string
	var results Results

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&results); err != nil {
		log.Fatal(err)
	}

	for _, result := range results.Results {
		mapLocations = append(mapLocations, result.Name)
	}

	return results.Config, mapLocations
}
