package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	semver    string
	commit    string
	buildDate string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display lld CLI version",
	Long:  "Display the version number of the Lemurian Labs device CLI",
	Run: func(cobra *cobra.Command, args []string) {
		fmt.Printf("lld CLI version %s (commit %s) (build Date: %s)\n", semver, commit, buildDate)
	},
}

// set by main.main() to marshal the semantic version and commit hash to this function
func SetVersion(semanticVersion string, commitHash string, bDate string) {
	if semanticVersion == "" {
		semver = "0.0.0"
	} else {
		semver = semanticVersion
	}

	if commitHash == "" {
		commit = "unknown"
	} else {
		commit = commitHash
	}

	if bDate == "" {
		buildDate = "unknown"
	} else {
		buildDate = bDate
	}
}
