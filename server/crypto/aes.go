package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"gategaurdian/server/constants"
	"gategaurdian/server/database/memorystore"
	"io"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 0o5}

// EncryptAes method is to encrypt or hide any classified text
func EncryptAes(text string) (string, error) {
	k, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvEncryptionKey)
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
	k, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvEncryptionKey)
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

// EncryptAesByte encrypt byte data using AES algorithm
func EncryptAesByte(text []byte) ([]byte, error) {
	var res []byte
	k, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvEncryptionKey)
	if err != nil {
		return res, err
	}
	key := []byte(k)
	c, err := aes.NewCipher(key)
	if err != nil {
		return res, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return res, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return res, err
	}

	return gcm.Seal(nonce, nonce, text, nil), nil
}

// DecryptAesByte decrypts byte data using AES algorithm
func DecryptAesByte(text []byte) ([]byte, error) {
	var res []byte
	k, err := memorystore.Provider.GetStringStoreEnvVariable(constants.EnvEncryptionKey)
	if err != nil {
		return res, err
	}
	key := []byte(k)
	c, err := aes.NewCipher(key)
	if err != nil {
		return res, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return res, err
	}

	nonceSize := gcm.NonceSize()
	if len(text) < nonceSize {
		return res, err
	}

	nonce, text := text[:nonceSize], text[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, text, nil)
	if err != nil {
		return res, err
	}

	return plaintext, nil
}
