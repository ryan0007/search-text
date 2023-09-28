package main

import (
	"log"

	"github.com/ryan0007/text-search/utils"
)

func main() {
	var textToSearch string = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
	var subtexts []string = []string{"Peter", "peter", "pick", "pi", "z", "Peterz", "me"}

	if e := utils.ValidateFields(textToSearch, subtexts); e != nil {
		log.Fatalln("Something went wrong:", e.Error())
	}

	for _, subtext := range subtexts {
		utils.PrintSubtextPositions(textToSearch, subtext)
	}
}
