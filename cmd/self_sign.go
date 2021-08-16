package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var selfSignCmd = &cobra.Command{
	Use:   "self-sign",
	Short: "Signs a CSR with it's own key.",
	Long:  "Signs a CSR with it's own key. This functionality is intended to be used to sign a Root CA's certificate.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return signCSR()
	},
}

func init() {
	selfSignCmd.Flags().StringVarP(&certificateFileName, "certFile", "c", "EmberCA.crt", "Where to save the resulting cert file.")
	selfSignCmd.Flags().StringVarP(&csrFileName, "csrFile", "f", "EmberCA.csr", "Path to the CSR to be signed.")
	selfSignCmd.Flags().StringVarP(&keyFileName, "keyFile", "k", "EmberCA.key", "Path to the key used to sign the CSR.")
	rootCmd.AddCommand(selfSignCmd)
}

func signCSR() error {
	log.Printf("Ding :)")
	return nil
}
