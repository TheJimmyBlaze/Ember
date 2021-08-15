package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thejimmyblaze/ember/internal/pki"
)

var csrFile string            //-f
var keyFile string            //-k
var subjectName string        //-n
var publicKeyAlgorithm string //-a
var publicKeyCurve string     //-c
var keyLength int             //-l

var csrCmd = &cobra.Command{
	Use:   "csr",
	Short: "Generates a CSR for Ember's Certificate Authority Certificate",
	Long:  "Generates a CA CSR, this may be self signed in the case of a root CA or signed by another CA in the case of an issuing or policy CA",
	RunE: func(cmd *cobra.Command, args []string) error {
		return pki.CreateCertificateSigningRequest(csrFile,
			keyFile,
			subjectName,
			publicKeyAlgorithm,
			publicKeyCurve,
			keyLength)
	},
}

func init() {
	csrCmd.Flags().StringVarP(&csrFile, "csrFile", "f", "EmberCA.csr", "Where to save the resulting CSR file")
	csrCmd.Flags().StringVarP(&keyFile, "keyFile", "k", "EmberCA.key", "Where to save the resulting Private Key file")
	csrCmd.Flags().StringVarP(&subjectName, "subjectName", "n", "CN=Ember-CertificateAuthority", "The Subject Name to describe the CA within the CSR")
	csrCmd.Flags().StringVarP(&publicKeyAlgorithm, "pkAlgo", "a", "RSA", `The assymetric encrypting algorithm used ("RSA", "ECDSA")`)
	csrCmd.Flags().StringVarP(&publicKeyCurve, "pkCurve", "c", "P384", `The elliptic curve to use for an ECDSA PK ("P224", "P256", "P384", "P512")`)
	csrCmd.Flags().IntVarP(&keyLength, "keyLength", "l", 4096, "The length of an RSA private key")
	rootCmd.AddCommand(csrCmd)
}
