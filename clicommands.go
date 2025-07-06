package main

import (
	"fmt"
	"github.com/daitomiun/gokedex/internal/service"
	. "github.com/daitomiun/gokedex/models"
	"os"
)

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
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

func commandMap(config *Config) error {
	newConfig, locations := service.GetMapLocations(*config.Next)
	//logic here to output the values
	fmt.Printf("config: %v locations: %v \n", newConfig, locations)
	fmt.Printf("config: %v locations: %v \n", newConfig.Next, locations)

	//logic here to update the config parameters after a new call

	return nil

}

func commandMapb(config *Config) error {
	newConfig, locations := service.GetMapLocations(*config.Prev)
	//logic here to output the values

	fmt.Printf("config: %v locations: %v \n", newConfig, locations)
	//logic here to update the config parameters after a new call
	// use the "next" and "prev" specs from the pokeAPI to show information about the maps
	return nil
}
