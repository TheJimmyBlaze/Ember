package cmd

import (
	"crypto/x509"
	"math/big"
	"time"

	"github.com/spf13/cobra"
	"github.com/thejimmyblaze/ember/pki"
)

//Some flags are used from root.go
var lifeSpanDays int //-d

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
	selfSignCmd.Flags().IntVarP(&lifeSpanDays, "lifespanDays", "d", 365, "Number of days before the self signed certificate will expire.")
	rootCmd.AddCommand(selfSignCmd)
}

func signCSR() error {

	//Load CSR
	csr, err := pki.LoadCSR(csrFileName)
	if err != nil {
		return err
	}

	//Load Key
	key, err := pki.LoadKey(keyFileName)
	if err != nil {
		return err
	}

	//Create certificate signing template from CSR
	template := &x509.Certificate{
		Signature:          csr.Signature,
		SignatureAlgorithm: csr.SignatureAlgorithm,

		PublicKeyAlgorithm: csr.PublicKeyAlgorithm,
		PublicKey:          csr.PublicKey,

		SerialNumber: big.NewInt(0),
		Issuer:       csr.Subject,
		Subject:      csr.Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Duration(lifeSpanDays) * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	//Sign CSR with Key
	signer := &pki.Signer{
		Cert: template,
		Key:  key,
	}
	cert, err := signer.SignCertificate(template, key.Public)
	if err != nil {
		return err
	}

	//Export
	err = cert.Export(certificateFileName)
	return err
}
