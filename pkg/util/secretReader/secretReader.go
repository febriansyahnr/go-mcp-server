package secret_reader

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
)

type PathCode string

const (
	PermataPathCode  PathCode = "permata"
	MockPathCode     PathCode = "mock"
	InternalPathCode PathCode = "snap_core"

	PRIVPKCS1 = "RSA PRIVATE KEY"
	PRIVPKCS8 = "PRIVATE KEY"

	PUBPKCS1 = "RSA PUBLIC KEY"
	PUBPKCS8 = "PUBLIC KEY"
)

type SecretReader interface {
	GetPublicKey() (*rsa.PublicKey, error)
	GetPrivateKey() (*rsa.PrivateKey, error)
	GetPgpPublicKey() ([]byte, error)
	GetPgpPrivateKey() ([]byte, error)
}

type pemReader struct {
	publicPath     string
	secretPath     string
	pgpPublicPath  string
	pgpPrivatePath string
}

func New(pathCode string) SecretReader {
	var p pemReader

	p.setPath(PathCode(strings.ToLower(pathCode)))

	return &p
}

func (p *pemReader) setPath(code PathCode) {
	folderPath := "./secrets-manager"
	p.publicPath = fmt.Sprintf("%s/%s_public.pem", folderPath, code)
	p.secretPath = fmt.Sprintf("%s/%s_private.pem", folderPath, code)
	p.pgpPublicPath = fmt.Sprintf("%s/%s_pgp_public.pem", folderPath, code)
	p.pgpPrivatePath = fmt.Sprintf("%s/%s_pgp_private.pem", folderPath, code)
}

func (p *pemReader) GetPublicKey() (*rsa.PublicKey, error) {

	pemData, err := os.ReadFile(p.publicPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	switch block.Type {
	case PUBPKCS8:
		p, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return p.(*rsa.PublicKey), err
	case PUBPKCS1:
		p, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return p, nil
	default:
		return nil, fmt.Errorf("public key not found")
	}
}

func (p *pemReader) GetPrivateKey() (*rsa.PrivateKey, error) {
	pemData, err := os.ReadFile(p.secretPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	switch block.Type {
	case PRIVPKCS8:
		p, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return p.(*rsa.PrivateKey), err
	case PRIVPKCS1:
		p, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return p, nil
	default:
		return nil, fmt.Errorf("private key not found")
	}
}

func (p *pemReader) GetPgpPublicKey() ([]byte, error) {
	pgpData, err := os.ReadFile(p.pgpPublicPath)
	if err != nil {
		return nil, err
	}

	return pgpData, nil
}

func (p *pemReader) GetPgpPrivateKey() ([]byte, error) {
	pgpData, err := os.ReadFile(p.pgpPrivatePath)
	if err != nil {
		return nil, err
	}

	return pgpData, nil
}
