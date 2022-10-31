package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func NewRSAKey(algo, keyID string) (*rsa.PrivateKey, string, string, string, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, "", "", "", err
	}

	privateKey, publicKey, err := AsRSAStr(key, &key.PublicKey)
	if err != nil {
		return nil, "", "", "", err
	}

	jwkPublicKey, err := GetPubJWK(algo, keyID, &key.PublicKey)
	if err != nil {
		return nil, "", "", "", err
	}

	return key, privateKey, publicKey, string(jwkPublicKey), err
}

func IsRSA(algo string) bool {
	switch algo {
	case "RS256", "RS384", "RS512":
		return true
	default:
		return false
	}
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkeyBytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkeyBytes,
		},
	)
	return string(privkeyPem)
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
	pubkeyBytes := x509.MarshalPKCS1PublicKey(pubkey)
	pubkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)

	return string(pubkeyPem)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub, nil
}

func AsRSAStr(privateKey *rsa.PrivateKey, publickKey *rsa.PublicKey) (string, string, error) {
	privPem := ExportRsaPrivateKeyAsPemStr(privateKey)
	pubPem := ExportRsaPublicKeyAsPemStr(publickKey)

	// Import the keys from pem string
	privParsed, err := ParseRsaPrivateKeyFromPemStr(privPem)
	if err != nil {
		return "", "", err
	}
	pubParsed, err := ParseRsaPublicKeyFromPemStr(pubPem)
	if err != nil {
		return "", "", err
	}

	privParsedPem := ExportRsaPrivateKeyAsPemStr(privParsed)
	pubParsedPem := ExportRsaPublicKeyAsPemStr(pubParsed)

	return privParsedPem, pubParsedPem, nil
}
