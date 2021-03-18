package misc

import (
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/base64"

	"github.com/ichi-infra-challenge/docker-local/api/src/logs"
)

//EncryptByGCM encrypt using GCM algorithm
func EncryptByGCM(plainKey []byte, plainText string) (string, error) {
	// make cipher block
	block, e := aes.NewCipher(plainKey)
	if e != nil {
		return "", e
	}
	// use gcm
	gcm, e := cipher.NewGCM(block)
	if e != nil {
		return "", e
	}
	// make nonce for gcm
	nonce := make([]byte, gcm.NonceSize())
	_, e = cryptoRand.Read(nonce)
	if e != nil {
		return "", e
	}
	// encrypt by gcm
	cipherByte := gcm.Seal(nil, nonce, []byte(plainText), nil)
	cipherByte = append(nonce, cipherByte...)
	cipherText := base64.StdEncoding.EncodeToString(cipherByte)
	return cipherText, nil
}

//DecryptByGCM decrypt using GCM algorithm
func DecryptByGCM(key []byte, cipherText string) (string, error) {
	// decode base64 string to bytes
	cipherBytes, e := base64.StdEncoding.DecodeString(cipherText)
	if e != nil {
		logs.Error("base64.StdEncoding.DecodeString(text)@DecryptByGCM", e, nil)
		return "", e
	}
	// make cipher block
	block, e := aes.NewCipher(key)
	if e != nil {
		return "", e
	}
	// use gcm
	gcm, e := cipher.NewGCM(block)
	if e != nil {
		return "", e
	}
	// get nonce for gcm
	nonce := cipherBytes[:gcm.NonceSize()]
	// decrypt by gcm
	plainByte, e := gcm.Open(nil, nonce, cipherBytes[gcm.NonceSize():], nil)
	if e != nil {
		return "", e
	}
	return string(plainByte), nil
}
