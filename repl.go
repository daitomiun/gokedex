package main

import (
	"bufio"
	"fmt"
	"github.com/daitomiun/gokedex/internal/pokecache"
	. "github.com/daitomiun/gokedex/models"
	"os"
	"strings"
	"time"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	output := strings.Fields(lower)
	return output
}

func startREPL() {
	reader := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(time.Duration(20 * time.Second))
	config := createConfig(20, 20)

	for true {
		fmt.Print("Pokedex > ")
		reader.Scan()
		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}
		cmd, exists := getAllCommands()[input[0]]
		if exists {
			err := cmd.Callback(&config, cache)
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

func createConfig(currentOffset, limit int32) Config {
	config := Config{
		Next:          currentOffset + limit,
		Prev:          currentOffset - limit,
		CurrentOffset: currentOffset,
		Limit:         limit,
	}
	if config.Prev < 0 {
		config.Prev = 0
	}
	return config
}
