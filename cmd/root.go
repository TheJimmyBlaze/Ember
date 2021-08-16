package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thejimmyblaze/ember/internal"
)

//There flags are shared among flags
var certificateFileName string //-c
var csrFileName string         //-f
var keyFileName string         //-k

var rootCmd = &cobra.Command{
	Use:   "ember-ca",
	Short: "Certificate Authority component of the lightweight Ember PKI solution.",
	Long: `Ember is a fast and lightweight PKI solution designed for deploying x.509 certificates
on a home lab or small business PKI. Ember CA provides the siging capability to the Ember system.
Ember CA uses a Certificate Authority certificate to issue, record, and revoke end-entity certificates.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
