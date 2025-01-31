package input

import (
	"strings"
)

func CleanInput(text string) []string {
	trimmedString := strings.Trim(text, " ")
	lowercased := strings.ToLower(trimmedString)
	cleanedInputs := strings.Fields(lowercased)
	return cleanedInputs
}
