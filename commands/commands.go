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
			description: "Displays a list of map areas",
			callback:    commandMap,
		},
	}
}
func getCommands() map[string]cliCommand {

	return commandsRegistry
}

func commandMap(cfg *config.Config) error {
	fmt.Printf("\n🗺️ Fetching Location Areas...\n\n")
	pokeClient := pokeapi.NewPokeClient()
	if cfg.Next == nil && cfg.Previous != nil {
		fmt.Printf("\n🚧 There are no further location areas to display\n\n")
		return nil
	}
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
	for _, location := range locationAreas.Results {
		fmt.Printf("📍 Location name: %s — with url: %s\n", location.Name, location.Url)
	}
	fmt.Println()
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
