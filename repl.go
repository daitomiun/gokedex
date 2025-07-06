package main

import (
	"bufio"
	"fmt"
	. "github.com/daitomiun/gokedex/models"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	output := strings.Fields(lower)
	return output
}

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	startUrl := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	config := Config{
		Next: &startUrl,
		Prev: nil,
	}
	fmt.Println(*config.Next)
	for true {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}
		cmd, exists := getAllCommands()[input[0]]
		if exists {
			err := cmd.Callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
