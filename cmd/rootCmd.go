package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const httpScheme = "http://"
const httpsScheme = "https://"

var rootCmd = &cobra.Command{
	Use:   "lld",
	Short: "Lemurian Labs Device cli",
	Long:  "CLI for the Lemurian Labs device simulators and services",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
