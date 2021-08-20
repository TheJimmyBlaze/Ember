package pki

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

const PEMBlockTypeCertificate string = "CERTIFICATE"

type Certificate struct {
	RawData []byte
	Cert    *x509.Certificate
}

func (cert *Certificate) Export(fileName string) error {

	log.Print("Converting Certificate to PEM format...")

	pemBlock := &pem.Block{
		Type:    PEMBlockTypeCertificate,
		Headers: nil,
		Bytes:   cert.RawData,
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("unable to create: %s for writing Certificate: %s", fileName, err)
	}
	defer file.Close()

	log.Printf("Writing Certificate to file: %s...", fileName)
	if err := pem.Encode(file, pemBlock); err != nil {
		return fmt.Errorf("unable to encode Certificate PEM block: %s", err)
	}
	return err
}
