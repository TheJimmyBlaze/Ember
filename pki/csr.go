package pki

import (
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

type CSR struct {
	RawData            []byte
	CertificateRequest *x509.CertificateRequest
	Key                *Key
}

func CreateCSR(subjectName string, publicKeyAlgorithm string, publicKeyCurve string, keyLength int) (*CSR, error) {

	log.Print("Creating Certificate Signing Request...")

	//Correctly case inputs
	publicKeyAlgorithm = strings.ToUpper(publicKeyAlgorithm)
	publicKeyCurve = strings.ToUpper(publicKeyCurve)

	//Generate Key
	key, err := CreateKey(publicKeyAlgorithm, publicKeyCurve, keyLength)
	if err != nil {
		return nil, err
	}

	//Generate CSR
	csr, err := generateCSR(subjectName, key)
	if err != nil {
		return nil, err
	}

	return csr, nil
}

func LoadCSR(fileName string) (csr *x509.CertificateRequest, err error) {

	//Open File
	pemBytes, err := os.ReadFile(fileName)
	if err != nil {
		return csr, fmt.Errorf("unable to open CSR file: %s, error: %s", fileName, err)
	}

	//Decode PEM to DER
	pemBlock, _ := pem.Decode(pemBytes)
	derBytes := pemBlock.Bytes

	//Parse
	csr, err = x509.ParseCertificateRequest(derBytes)
	if err != nil {
		return csr, fmt.Errorf("unable to parse CSR from file data: %s, error: %s", fileName, err)
	}

	return csr, nil
}

func generateCSR(subjectName string, key *Key) (*CSR, error) {

	log.Print("Creating CSR...")

	//Convert Subject Name
	log.Printf("Converting Subject Name: %s to PKIX Name...", subjectName)
	subject, err := x500dn.ParseDN(subjectName)
	if err != nil {
		return nil, fmt.Errorf("could not create CSR template: subject Name invalid, ensure it is in x500 format. Error: %s", err)
	}

	//Build Request
	template := &x509.CertificateRequest{
		Subject: *subject,
	}

	request, err := x509.CreateCertificateRequest(rand.Reader, template, key.Private)
	if err != nil {
		return nil, fmt.Errorf("could not create CSR: %s", err)
	}

	//Parse request back into Certificate Request struct
	certificateRequest, err := x509.ParseCertificateRequest(request)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal CSR back into struct: %s", err)
	}

	//Return struct
	csr := &CSR{
		RawData:            request,
		CertificateRequest: certificateRequest,
		Key:                key,
	}

	return csr, nil
}

func (csr *CSR) Export(fileName string) error {

	log.Print("Converting CSR to PEM format...")

	pemBlock := &pem.Block{
		Type:    PEMBlockTypeCSR,
		Headers: nil,
		Bytes:   csr.RawData,
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
