package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexdevdavis/pokedexcli/input"
	"github.com/alexdevdavis/pokedexcli/commands"
	"github.com/alexdevdavis/pokedexcli/config"

)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config.Config{}
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			dirtyLine := scanner.Text()
			line := input.CleanInput(dirtyLine)
			err := commands.ExecuteCommand(line[0], cfg)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}
