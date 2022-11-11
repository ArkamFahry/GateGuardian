package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// NewECDSAKey to generate new ECDSA Key if env is not set
// returns key instance, private key string, public key string, jwk string, error
func NewECDSAKey(algo, keyID string) (*ecdsa.PrivateKey, string, string, string, error) {
	var curve elliptic.Curve
	switch algo {
	case "ES256":
		curve = elliptic.P256()
	case "ES384":
		curve = elliptic.P384()
	case "ES512":
		curve = elliptic.P521()
	default:
		return nil, "", "", "", errors.New("invalid algo")
	}
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, "", "", "", err
	}

	privateKey, publicKey, err := AsECDSAStr(key, &key.PublicKey)
	if err != nil {
		return nil, "", "", "", err
	}

	jwkPublicKey, err := GetPubJWK(algo, keyID, &key.PublicKey)
	if err != nil {
		return nil, "", "", "", err
	}

	return key, privateKey, publicKey, string(jwkPublicKey), err
}

// IsECDSA checks if given string is valid ECDSA algo
func IsECDSA(algo string) bool {
	switch algo {
	case "ES256", "ES384", "ES512":
		return true
	default:
		return false
	}
}

// ExportEcdsaPrivateKeyAsPemStr to get ECDSA private key as pem string
func ExportEcdsaPrivateKeyAsPemStr(privatekey *ecdsa.PrivateKey) (string, error) {
	privateKeyBytes, err := x509.MarshalECPrivateKey(privatekey)
	if err != nil {
		return "", err
	}
	privateKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "ECDSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)
	return string(privateKeyPem), nil
}

// ExportEcdsaPublicKeyAsPemStr to get ECDSA public key as pem string
func ExportEcdsaPublicKeyAsPemStr(publicKey *ecdsa.PublicKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	publicKeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "ECDSA PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)

	return string(publicKeyPem), nil
}

// ParseEcdsaPrivateKeyFromPemStr to parse ECDSA private key from pem string
func ParseEcdsaPrivateKeyFromPemStr(privPEM string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

// ParseEcdsaPublicKeyFromPemStr to parse ECDSA public key from pem string
func ParseEcdsaPublicKeyFromPemStr(pubPEM string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *ecdsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("key type is not ECDSA")
}

// AsECDSAStr returns private, public key string or error
func AsECDSAStr(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string, error) {
	// Export the keys to pem string
	privPem, err := ExportEcdsaPrivateKeyAsPemStr(privateKey)
	if err != nil {
		return "", "", err
	}
	pubPem, err := ExportEcdsaPublicKeyAsPemStr(publicKey)
	if err != nil {
		return "", "", err
	}

	// Import the keys from pem string
	privParsed, err := ParseEcdsaPrivateKeyFromPemStr(privPem)
	if err != nil {
		return "", "", err
	}
	pubParsed, err := ParseEcdsaPublicKeyFromPemStr(pubPem)
	if err != nil {
		return "", "", err
	}

	// Export the newly imported keys
	privParsedPem, err := ExportEcdsaPrivateKeyAsPemStr(privParsed)
	if err != nil {
		return "", "", err
	}
	pubParsedPem, err := ExportEcdsaPublicKeyAsPemStr(pubParsed)
	if err != nil {
		return "", "", err
	}

	return privParsedPem, pubParsedPem, nil
}
