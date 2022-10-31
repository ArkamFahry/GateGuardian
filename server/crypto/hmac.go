package crypto

import (
	"github.com/google/uuid"
)

func NewHMACKey(algo, keyID string) (string, string, error) {
	key := uuid.New().String()
	jwkPublicKey, err := GetPubJWK(algo, keyID, []byte(key))
	if err != nil {
		return "", "", err
	}
	return key, string(jwkPublicKey), nil
}

func IsHMACA(algo string) bool {
	switch algo {
	case "HS256", "HS384", "HS512":
		return true
	default:
		return false
	}
}
