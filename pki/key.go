package pki

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
)

const PKAlgorithmRSA string = "RSA"
const PKAlgorithmECDSA string = "ECDSA"

const PEMBlockTypePK string = "PRIVATE KEY"

var ellipticCurveDetails = []struct {
	curve elliptic.Curve
	name  string
}{
	{elliptic.P224(), "P225"},
	{elliptic.P256(), "P256"},
	{elliptic.P384(), "P384"},
	{elliptic.P521(), "P521"},
}

type Key struct {
	Private crypto.Signer
	Public  crypto.PublicKey
}

func CreateKey(publicKeyAlgorithm string, publicKeyCurve string, keyLength int) (*Key, error) {

	log.Print("Generating CSR key pair...")

	switch publicKeyAlgorithm {

	case PKAlgorithmRSA:
		log.Print("Creating RSA Private Key...")
		privateKey, err := rsa.GenerateKey(rand.Reader, keyLength)
		if err != nil {
			return nil, err
		}
		key := &Key{
			Private: privateKey,
			Public:  &privateKey.PublicKey,
		}
		return key, err

	case PKAlgorithmECDSA:
		log.Print("Creating ECDSA Private Key...")
		curve, err := determineCurve(publicKeyCurve)
		if err != nil {
			return nil, err
		}
		privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			return nil, err
		}
		key := &Key{
			Private: privateKey,
			Public:  &privateKey.PublicKey,
		}
		return key, err
	}

	return nil, errors.New("key pair creation failed, no public key algorithm was available")
}

func LoadKey(fileName string) (*Key, error) {

	//Open File
	pemBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to open PK file: %s, error: %s", fileName, err)
	}

	//Convert PEM to DER
	pemBlock, _ := pem.Decode(pemBytes)
	derBytes := pemBlock.Bytes

	//Parse key
	keyInterface, err := x509.ParsePKCS8PrivateKey(derBytes)
	if err != nil {
		return nil, fmt.Errorf("unable to parse PKCS8 key from file: %s, error: %s", fileName, err)
	}

	//Cast to crypto.Signer
	switch privateKey := keyInterface.(type) {
	case *rsa.PrivateKey:
		key := &Key{
			Private: privateKey,
			Public:  &privateKey.PublicKey,
		}
		return key, nil
	case *ecdsa.PrivateKey:
		key := &Key{
			Private: privateKey,
			Public:  &privateKey.PublicKey,
		}
		return key, nil
	default:
		return nil, fmt.Errorf("PK not in recognised format, key must be a PKCS8 formatted RSA or ECDSA key")
	}
}

func (key *Key) Export(fileName string) error {

	log.Print("Converting PK to PEM format...")

	//Create PEM
	bytes, err := x509.MarshalPKCS8PrivateKey(key.Private)
	if err != nil {
		return fmt.Errorf("unable to marshal private key to PKCS8 format: %s", err)
	}
	pemBlock := &pem.Block{
		Type:    PEMBlockTypePK,
		Headers: nil,
		Bytes:   bytes,
	}

	//Save PEM to file
	log.Printf("Writing PK to file: %s...", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("unable to create: %s for writing PK: %s", fileName, err)
	}
	defer file.Close()

	if err := pem.Encode(file, pemBlock); err != nil {
		return fmt.Errorf("unable to encode PK PEM block: %s", err)
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
