package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thejimmyblaze/ember/ember"
	"github.com/thejimmyblaze/ember/version"
)

var rootCmd = &cobra.Command{
	Use:   "ember-ca",
	Short: "Certificate Authority component of the lightweight Ember PKI solution",
	Long: `Ember is a fast and lightweight PKI solution designed for deploying x.509 certificates
on a home lab or small business PKI. Ember CA provides the siging capability to the Ember system.
Ember CA uses a Certificate Authority certificate to issue, record, and revoke end-entity certificates`,
	Run: func(cmd *cobra.Command, args []string) {
		ember.Launch()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version information of Ember CA",
	Long:  "Prints the build version, when the build was created, and the hash of the build files",
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintVersionInformation()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
