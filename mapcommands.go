package main

import (
	. "github.com/daitomiun/gokedex/models"
)

func getAllCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "explore the list of available pokemons in an area",
			Callback:    commandExplore,
		},
	}
}
