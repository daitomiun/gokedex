package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/daitomiun/gokedex/internal/pokecache"
	. "github.com/daitomiun/gokedex/models"
	"io"
	"log"
	"net/http"
)

func GetMapLocations(offset int32, cache *pokecache.Cache) []string {
	url := createUrl(offset)
	println(url)
	entry, exists := cache.Get(url)
	if exists {
		bytesReader := bytes.NewReader(entry)
		return getLocations(io.NopCloser(bytesReader))
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	cachedBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	cache.Add(url, cachedBody)
	bytesReader := bytes.NewReader(cachedBody)
	return getLocations(io.NopCloser(bytesReader))
}

func getLocations(entry io.ReadCloser) []string {
	var mapLocations []string
	var results Results

	decoder := json.NewDecoder(entry)
	if err := decoder.Decode(&results); err != nil {
		log.Fatal(err)
	}

	for _, result := range results.Results {
		mapLocations = append(mapLocations, result.Name)
	}
	return mapLocations
}

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

func createUrl(offset int32) string {
	return fmt.Sprintf("%s?offset=%v&limit=20", baseUrl, offset)
}
