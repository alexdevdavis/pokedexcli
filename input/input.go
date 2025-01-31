package input

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	trimmedString := strings.Trim(text, " ")
	lowercased := strings.ToLower(trimmedString)
	cleanedInputs := strings.Fields(lowercased)
	fmt.Println(cleanedInputs)
	return cleanedInputs
}
