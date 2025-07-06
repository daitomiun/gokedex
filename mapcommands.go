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
			Description: "Shows next pokemon 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Shows previous pokemon 20 locations",
			Callback:    commandMapb,
		},
	}
}
