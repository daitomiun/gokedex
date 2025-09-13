package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/daitomiun/gokedex/internal/pokecache"
	. "github.com/daitomiun/gokedex/models"
)

func GetPokemonsFromLocation(location string, cache *pokecache.Cache) []string {
	url := setLocationUrl(location)
	println(url)
	entry, exists := cache.Get(url)
	if exists {
		bytesReader := bytes.NewReader(entry)
		return getPokemons(io.NopCloser(bytesReader))
	}
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	cachedBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	cache.Add(url, cachedBody)
	bytesReader := bytes.NewReader(cachedBody)
	return getPokemons(io.NopCloser(bytesReader))
}

func GetMapLocations(offset int32, cache *pokecache.Cache) []string {
	url := setOffsetUrl(offset)
	println(url)
	entry, exists := cache.Get(url)
	if exists {
		bytesReader := bytes.NewReader(entry)
		return getLocations(io.NopCloser(bytesReader))
	}
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	cachedBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
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

func getPokemons(entry io.ReadCloser) []string {
	var pokemons []string
	var results Pokemons

	decoder := json.NewDecoder(entry)
	if err := decoder.Decode(&results); err != nil {
		log.Fatal(err)
	}

	for _, result := range results.PokemonEncounters {
		pokemons = append(pokemons, result.Pokemon.Name)
	}
	return pokemons
}

const baseUrl = "https://pokeapi.co/api/v2/location-area/"

func setOffsetUrl(offset int32) string {
	return fmt.Sprintf("%s?offset=%v&limit=20", baseUrl, offset)
}

func setLocationUrl(location string) string {
	return baseUrl + location
}
