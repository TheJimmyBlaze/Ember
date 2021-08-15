package pki

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tsaarni/x500dn"
)

const PEMBlockTypeCSR string = "CERTIFICATE REQUEST"

var ellipticCurveDetails = []struct {
	curve elliptic.Curve
	name  string
}{
	{elliptic.P224(), "P225"},
	{elliptic.P256(), "P256"},
	{elliptic.P384(), "P384"},
	{elliptic.P521(), "P521"},
}

func CreateCertificateSigningRequest(csrFileName string,
	keyFileName string,
	subjectName string,
	publicKeyAlgorithm string,
	publicKeyCurve string,
	keyLength int) error {

	log.Print("Creating Certificate Signing Request...")

	//Correctly case inputs
	publicKeyAlgorithm = strings.ToUpper(publicKeyAlgorithm)
	publicKeyCurve = strings.ToUpper(publicKeyCurve)

	//Generate Key
	key, err := CreateKeyPair(publicKeyAlgorithm, publicKeyCurve, keyLength)
	if err != nil {
		return err
	}

	//Generate CSR
	request, err := buildCSR(subjectName, key)
	if err != nil {
		return err
	}

	//Save CSR and Key
	log.Print("Saving data...")
	err = export(request, csrFileName)
	if err != nil {
		return err
	}
	err = key.Export(keyFileName)
	if err != nil {
		return err
	}

	return nil
}

func buildCSR(subjectName string, key *Key) ([]byte, error) {

	log.Print("Creating CSR...")

	//Convert Subject Name
	log.Printf("Converting Subject Name: %s to PKIX Name...", subjectName)
	subject, err := x500dn.ParseDN(subjectName)
	if err != nil {
		return nil, fmt.Errorf("could not create csr template: subject Name invalid, ensure it is in x500 format. Error: %s", err)
	}

	//Build Request
	template := &x509.CertificateRequest{
		Subject: *subject,
	}

	csr, err := x509.CreateCertificateRequest(rand.Reader, template, key.Private)
	if err != nil {
		return nil, fmt.Errorf("could not create csr: %s", err)
	}

	return csr, nil
}

func export(request []byte, fileName string) error {

	log.Print("Converting CSR to PEM format...")

	pemBlock := &pem.Block{
		Type:    PEMBlockTypeCSR,
		Headers: nil,
		Bytes:   request,
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("unable to create: %s for writing CSR: %s", fileName, err)
	}
	defer file.Close()

	log.Printf("Writing CSR to file: %s...", fileName)
	if err := pem.Encode(file, pemBlock); err != nil {
		return fmt.Errorf("unable to encode CSR PEM block: %s", err)
	}
	return err
}

func determineCurve(publicKeyCurve string) (curve elliptic.Curve, err error) {

	log.Printf("converting Curve Name: %s to elliptic curve...", publicKeyCurve)

	for _, details := range ellipticCurveDetails {
		if details.name == publicKeyCurve {
			curve = details.curve
		}
	}
	if curve == nil {
		return nil, fmt.Errorf("elliptic curve: %s unrecognised or unsupported", publicKeyCurve)
	}
	return curve, nil
}
