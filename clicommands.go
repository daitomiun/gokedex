package main

import (
	"fmt"
	"github.com/daitomiun/gokedex/internal/pokecache"
	"github.com/daitomiun/gokedex/internal/service"
	. "github.com/daitomiun/gokedex/models"
	"os"
)

func commandExit(config *Config, cache *pokecache.Cache, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, cache *pokecache.Cache, param string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, cmd := range getAllCommands() {
		fmt.Printf("%s: %s \n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *Config, cache *pokecache.Cache, param string) error {
	locations := service.GetMapLocations(config.CurrentOffset, cache)
	if len(locations) == 0 {
		fmt.Println("End of next locations try to go back")
		return nil
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	config.CurrentOffset += config.Limit
	config.Next = config.CurrentOffset + config.Limit
	config.Prev = config.CurrentOffset - config.Limit
	if config.Prev < 0 {
		config.Prev = 0
	}
	return nil
}

func commandMapb(config *Config, cache *pokecache.Cache, param string) error {
	if config.CurrentOffset == 0 {
		fmt.Println("You're on the first page")
		return nil
	}
	config.CurrentOffset -= config.Limit
	if config.CurrentOffset < 0 {
		config.CurrentOffset = 0
	}
	locations := service.GetMapLocations(config.CurrentOffset, cache)
	if len(locations) == 0 || locations == nil {
		fmt.Println("No locations found")
		return nil
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	config.Next = config.CurrentOffset + config.Limit
	config.Prev = config.CurrentOffset - config.Limit
	if config.Prev < 0 {
		config.Prev = 0
	}
	return nil
}

func commandExplore(config *Config, cache *pokecache.Cache, location string) error {
	pokemons := service.GetPokemonsFromLocation(location, cache)
	fmt.Printf("Exploring %s \n", location)
	if len(pokemons) == 0 || pokemons == nil {
		fmt.Printf("No pokemon on %s ", location)
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons {
		fmt.Printf("- %s \n", pokemon)
	}
	return nil
}
