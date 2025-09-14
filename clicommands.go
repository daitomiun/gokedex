package main

import (
	"fmt"
	"github.com/daitomiun/gokedex/internal/pokecache"
	"github.com/daitomiun/gokedex/internal/service"
	. "github.com/daitomiun/gokedex/models"
	"math/rand"
	"os"
)

func commandExit(config *Config, cache *pokecache.Cache, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, cache *pokecache.Cache, _ string) error {
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

func commandMap(config *Config, cache *pokecache.Cache, _ string) error {
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

func commandMapb(config *Config, cache *pokecache.Cache, _ string) error {
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

func commandCatch(config *Config, cache *pokecache.Cache, pokemon string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	entry, exists := config.Pokedex[pokemon]
	if exists {
		fmt.Printf("You already have %s! \n", entry.Name)
		return nil
	}
	newPokemon := service.GetPokemon(pokemon)
	if len(newPokemon.Name) == 0 {
		fmt.Printf("Pokemon %s not found! \n", pokemon)
		return nil
	}
	probability := 0.0
	if newPokemon.BaseExperience >= 0 && newPokemon.BaseExperience <= 40 {
		probability = 0.8
	} else if newPokemon.BaseExperience > 40 && newPokemon.BaseExperience <= 200 {
		probability = 0.5
	} else {
		probability = 0.2
	}
	caught := rand.Float64() < probability
	if caught {
		config.Pokedex[newPokemon.Name] = newPokemon
		fmt.Printf("%s was caught! \n", newPokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s was not caught, try again! \n", newPokemon.Name)
	}
	return nil
}

func commandInspect(config *Config, cache *pokecache.Cache, pokemon string) error {
	entry, exists := config.Pokedex[pokemon]
	if !exists {
		fmt.Printf("You don't have %s! \n", pokemon)
		return nil
	}
	fmt.Printf(`
Name: %s
Height: %d
Weight: %d 
`,
		entry.Name,
		entry.Height,
		entry.Weight,
	)
	fmt.Println("Stats:")
	for _, elem := range entry.Stats {
		fmt.Printf("  -%s: %d\n", elem.Stat.Name, elem.BaseStat)
	}
	fmt.Println("Types:")
	for _, elem := range entry.Types {
		fmt.Printf("  - %s\n", elem.Type.Name)
	}
	return nil
}

func commandPokedex(config *Config, cache *pokecache.Cache, pokemon string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
