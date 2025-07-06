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
	if config.Next == nil {
		fmt.Println("End of next locations try to go back")
		return nil
	}
	newConfig, locations := service.GetMapLocations(*config.Next)
	*config = newConfig
	for _, location := range locations {
		fmt.Println(location)
	}

	return nil
}

func commandMapb(config *Config) error {
	if config.Prev == nil {
		fmt.Println("End of Prev locations try to go next")
		return nil
	}
	newConfig, locations := service.GetMapLocations(*config.Prev)
	*config = newConfig
	for _, location := range locations {
		fmt.Println(location)
	}

	return nil
}
