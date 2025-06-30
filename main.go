package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(reader)
		scanner.Scan()

		cmd := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %v \n", cmd[0])
	}
}
