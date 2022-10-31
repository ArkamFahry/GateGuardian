package crypto

import "encoding/base64"

func EncryptB64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func DecryptB64(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
