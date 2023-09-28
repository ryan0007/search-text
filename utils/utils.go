package utils

import (
	"fmt"
	"strings"
)

func findSubtextIndexes(textToSearch, subtext string) []int {
	textToSearch = strings.ToLower(textToSearch)
	subtext = strings.ToLower(subtext)
	subtextLength := len(subtext)
	positions := make([]int, 0)

	for index := 0; index <= len(textToSearch)-subtextLength; index++ {
		if textToSearch[index:index+subtextLength] == subtext {
			positions = append(positions, index+1)
		}
	}

	return positions
}

func PrintSubtextPositions(textToSearch, subtext string) {
	positions := findSubtextIndexes(textToSearch, subtext)

	if len(positions) > 0 {
		fmt.Printf("%s \"%v\"\n", subtext, positions)
	} else {
		fmt.Printf("%s \"<No Output>\"\n", subtext)
	}
}

func ValidateFields(textToSearch string, subtexts []string) error {
	if textToSearch == "" || len(subtexts) == 0 {
		return fmt.Errorf("textToSearch or subtext can't be empty")
	}

	return nil
}
