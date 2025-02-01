package commands

import (
	"fmt"
	"os"

	"github.com/alexdevdavis/pokedexcli/config"
	"github.com/alexdevdavis/pokedexcli/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config) error
}

func ExecuteCommand(command string, cfg *config.Config) error {
	cmd, exists := getCommands()[command]
	if !exists {
		return fmt.Errorf("Unknown command")
	}
	err := cmd.callback(cfg)
	return err
}

var commandsRegistry map[string]cliCommand

func init() {
	commandsRegistry = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of 20 location areas",
			callback:    commandMapb,
		},
	}
}
func getCommands() map[string]cliCommand {

	return commandsRegistry
}

func commandMapb(cfg *config.Config) error {
	if cfg.Previous == nil {
		fmt.Printf("\n\nYou're on the first page\n\n")
		return nil
	}
	pokeClient := pokeapi.NewPokeClient()
	locationAreas, err := pokeClient.LocationAreas(cfg.Previous)
	if err != nil {
		return err
	}
	if len(locationAreas.Results) == 0 {
		fmt.Printf("\nThere are no further location areas to display\n\n")
		return nil
	}
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous
	printLocationData(locationAreas.Results)
	return nil

}

func commandMap(cfg *config.Config) error {
	if cfg.Next == nil && cfg.Previous != nil {
		fmt.Printf("\n🚧 There are no further location areas to display\n\n")
		return nil
	}
	fmt.Printf("\n🗺️ Fetching Location Areas...\n\n")
	pokeClient := pokeapi.NewPokeClient()

	locationAreas, err := pokeClient.LocationAreas(cfg.Next)
	if err != nil {
		return err
	}
	if len(locationAreas.Results) == 0 {
		fmt.Printf("\nThere are no further location areas to display\n\n")
		return nil
	}
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous
	printLocationData(locationAreas.Results)
	return nil
}

func commandHelp(cfg *config.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	commands := getCommands()
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config.Config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func printLocationData(locationAreas []pokeapi.LocationArea) {
	for _, location := range locationAreas {
		fmt.Printf("📍 Location name: %s — with url: %s\n", location.Name, location.Url)
	}
	fmt.Println()

}
