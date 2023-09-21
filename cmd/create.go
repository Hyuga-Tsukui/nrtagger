package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/Hyuga-Tsukui/nrtagger/cmd/internal/utils"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new deployment tag",
	Run: func(cmd *cobra.Command, args []string) {

		aliases := utils.ReadAliases()

		var (
			guid    string
			profile string
		)

		if len(aliases) > 0 {
			alias := utils.SelectAlias(aliases)
			guid = alias.GUID
			profile = alias.Profile
		} else {
			guid = utils.Prompt("GUID", "")
			profile = utils.SelectProfile()
		}
		version := utils.Prompt("Version", "")
		commitHash := utils.Prompt("Commit Hash", getDefaultCommitHash())
		description := utils.Prompt("Description", "")

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
