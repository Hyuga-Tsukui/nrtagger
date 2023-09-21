package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

type ProfileCredentials map[string]interface{}

func readProfileCredentials() ProfileCredentials {
	filePath := filepath.Join(os.Getenv("HOME"), ".newrelic", "credentials.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("\033[1;31mError: credentials file does not exist. Please set up the New Relic CLI.\033[0m\n")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error: Unable to read profile credentials from %s, %v\n", filePath, err)
	}

	var credentials ProfileCredentials
	err = json.Unmarshal(content, &credentials)
	if err != nil {
		log.Fatalf("Error: Unable to parse profile credentials from %s, %v\n", filePath, err)
	}

	return credentials
}

func SelectProfile() string {
	credentials := readProfileCredentials()
	keys := make([]string, 0, len(credentials))
	for key := range credentials {
		keys = append(keys, key)
	}

	prompt := promptui.Select{
		Label: "Select Profile",
		Items: keys,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Selection failed %v\n", err)
	}

	return result
}
