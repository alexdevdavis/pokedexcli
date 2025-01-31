package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func ExecuteCommand(command string) error {
	cmd, exists := getCommands()[command]
	if !exists {
		return fmt.Errorf("Unknown command")
	}
	err := cmd.callback()
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
	}
}
func getCommands() map[string]cliCommand {

	return commandsRegistry
}

func commandHelp() error {
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

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}
