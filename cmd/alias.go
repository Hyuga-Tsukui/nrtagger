package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Hyuga-Tsukui/nrtagger/cmd/internal/utils"
	"github.com/spf13/cobra"
)

const (
	guidKey    = "guid"
	profileKey = "profile"
)

var aliasFile = filepath.Join(os.Getenv("HOME"), ".nrtagger/alias.json")

var aliasCmd = &cobra.Command{
	Use:   "add-alias",
	Short: "Create alias for GUIDs & Profiles",
	RunE:  runAlias,
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}

func runAlias(cmd *cobra.Command, args []string) error {
	aliasName := utils.Prompt("Enter alias name", "")
	profile := utils.SelectProfile()
	guid := utils.Prompt("Enter GUID", "")

	aliases := make(map[string]map[string]string)

	if _, err := os.Stat(aliasFile); !os.IsNotExist(err) {
		content, err := os.ReadFile(aliasFile)
		if err != nil {
			return err
		}

		err = json.Unmarshal(content, &aliases)
		if err != nil {
			return err
		}
	}

	if _, ok := aliases[aliasName]; !ok {
		aliases[aliasName] = make(map[string]string)
	}

	aliases[aliasName][guidKey] = guid
	aliases[aliasName][profileKey] = profile

	content, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Dir(aliasFile)); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(aliasFile), 0755)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(aliasFile, content, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Alias added successfully!")

	return nil
}
