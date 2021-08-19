package pki

import "crypto/x509"

type Signer struct {
	Key *Key
}

func (signer *Signer) SignCSR(csr *x509.CertificateRequest) (*x509.Certificate, error) {
	return nil, nil
}
