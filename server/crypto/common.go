package crypto

import (
	"crypto/x509"

	"github.com/ArkamFahry/GateGuardian/server/constants"
	"github.com/ArkamFahry/GateGuardian/server/memorystore/envstore"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/square/go-jose.v2"
)

// GetPubJWK returns JWK for given keys
func GetPubJWK(algo, keyID string, publicKey interface{}) (string, error) {
	jwk := &jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{
				Algorithm:                   algo,
				Key:                         publicKey,
				Use:                         "sig",
				KeyID:                       keyID,
				Certificates:                []*x509.Certificate{},
				CertificateThumbprintSHA1:   []uint8{},
				CertificateThumbprintSHA256: []uint8{},
			},
		},
	}
	jwkPublicKey, err := jwk.Keys[0].MarshalJSON()
	if err != nil {
		return "", err
	}
	return string(jwkPublicKey), nil
}

func GenerateJWKBasedOnEnv() (string, error) {
	jwk := ""
	algo, err := envstore.Provider.GetEnv(constants.JWT_TYPE)
	if err != nil {
		return jwk, err
	}
	clientID, err := envstore.Provider.GetEnv(constants.CLIENT_ID)
	if err != nil {
		return jwk, err
	}

	jwtSecret, err := envstore.Provider.GetEnv(constants.JWT_SECRET)
	if err != nil {
		return jwk, err
	}

	// check if jwt secret is provided
	if IsHMACA(algo) {
		jwk, err = GetPubJWK(algo, clientID, []byte(jwtSecret))
		if err != nil {
			return "", err
		}
	}

	jwtPublicKey, err := envstore.Provider.GetEnv(constants.JWT_PUBLIC_KEY)
	if err != nil {
		return jwk, err
	}

	if IsRSA(algo) {
		publicKeyInstance, err := ParseRsaPublicKeyFromPemStr(jwtPublicKey)
		if err != nil {
			return "", err
		}

		jwk, err = GetPubJWK(algo, clientID, publicKeyInstance)
		if err != nil {
			return "", err
		}
	}

	if IsECDSA(algo) {
		jwtPublicKey, err = envstore.Provider.GetEnv(constants.JWT_PUBLIC_KEY)
		if err != nil {
			return jwk, err
		}
		publicKeyInstance, err := ParseEcdsaPublicKeyFromPemStr(jwtPublicKey)
		if err != nil {
			return "", err
		}

		jwk, err = GetPubJWK(algo, clientID, publicKeyInstance)
		if err != nil {
			return "", err
		}
	}

	return jwk, nil
}

func EncryptPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pw), nil
}
