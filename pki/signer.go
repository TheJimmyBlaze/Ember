package pki

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
)

type Signer struct {
	Cert *x509.Certificate
	Key  *Key
}

func (signer *Signer) SignCertificate(csr *x509.Certificate, publicKey crypto.PublicKey) (cert Certificate, err error) {

	log.Printf("Signing Certificate: %s...", csr.Subject.CommonName)

	//Convert public key
	var realKey interface{}
	realKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		realKey, ok = publicKey.(*ecdsa.PublicKey)
		if !ok {
			return cert, fmt.Errorf("write this")
		}
	}

	//Sign
	certBytes, err := x509.CreateCertificate(rand.Reader, csr, signer.Cert, realKey, signer.Key.Private)
	if err != nil {
		return cert, fmt.Errorf("unable to sign certificate: %s", err)
	}

	cert = Certificate{
		RawData: certBytes,
		Cert:    nil,
	}

	return cert, nil
}
