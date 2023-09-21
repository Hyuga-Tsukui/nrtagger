package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

type Aliases map[string]map[string]string

type Alias struct {
	Name    string
	GUID    string
	Profile string
}

func SelectAlias(aliases Aliases) Alias {
	keys := make([]string, 0, len(aliases))
	for key := range aliases {
		keys = append(keys, key)
	}

	prompt := promptui.Select{
		Label: "Select Alias",
		Items: keys,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection failed %v\n", err)
	}

	return Alias{
		Name:    result,
		GUID:    aliases[result]["guid"],
		Profile: aliases[result]["profile"],
	}
}

func ReadAliases() Aliases {
	filePath := filepath.Join(os.Getenv("HOME"), ".nrtagger", "alias.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return make(Aliases, 0)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: Unable to read aliases from %s, %v\n", filePath, err)
	}

	var aliases Aliases
	err = json.Unmarshal(content, &aliases)
	if err != nil {
		log.Fatalf("Error: Unable to parse aliases from %s, %v\n", filePath, err)
	}

	return aliases
}
