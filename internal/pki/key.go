package pki

import (
	"crypto"
	"crypto/ecdsa"
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

type Key struct {
	Private crypto.PrivateKey
	Public  crypto.PublicKey
}

func CreateKeyPair(publicKeyAlgorithm string, publicKeyCurve string, keyLength int) (*Key, error) {

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
