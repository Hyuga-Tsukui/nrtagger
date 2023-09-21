package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new deployment tag",
	Run: func(cmd *cobra.Command, args []string) {
		profile := selectProfile()
		guid := prompt("GUID", "")
		version := prompt("Version", "")
		commitHash := prompt("Commit Hash", getDefaultCommitHash())
		description := prompt("Description", "")

		command := fmt.Sprintf(
			"newrelic --profile %s entity deployment create --guid %s --version %s --commit %s --description %s",
			profile, guid, version, commitHash, description,
		)

		err := exec.Command("sh", "-c", command).Run()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func prompt(label string, defaultValue string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result
}

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

func selectProfile() string {
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

func getDefaultCommitHash() string {
	cmd := exec.Command("git", "show", "--format='%H'", "--no-patch")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf("\033[1;33mWarning: Unable to get default commit hash, ensure you are in a git managed directory. Error: %v\033[0m\n", err)
		return ""
	}
	return strings.TrimSpace(out.String())
}
