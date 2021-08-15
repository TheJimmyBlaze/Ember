package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thejimmyblaze/ember/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version information of Ember CA",
	Long:  "Prints the build version, when the build was created, and the hash of the build files",
	Run: func(cmd *cobra.Command, args []string) {
		printVersionInformation()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printVersionInformation() {
	fmt.Printf("Ember - X.509 Crypto Service - %s\n", version.BuildVersion)
	fmt.Printf("Build Time: %s\n", version.BuildTime)
	fmt.Printf("Build Hash: %s\n", version.BuildHash)
}
