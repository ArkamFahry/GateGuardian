package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"gategaurdian/server/constants"
	"gategaurdian/server/database/memorydb"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 0o5}

// EncryptAes method is to encrypt or hide any classified text
func EncryptAes(text string) (string, error) {
	k, err := memorydb.Provider.GetStringStoreEnvVariable(constants.EnvEncryptionKey)
	if err != nil {
		return "", err
	}
	key := []byte(k)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return EncryptB64(string(cipherText)), nil
}

// DecryptAes method is to extract back the encrypted text
func DecryptAes(text string) (string, error) {
	k, err := memorydb.Provider.GetStringStoreEnvVariable(constants.EnvEncryptionKey)
	if err != nil {
		return "", err
	}
	key := []byte(k)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipherText, err := DecryptB64(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, []byte(cipherText))
	return string(plainText), nil
}
