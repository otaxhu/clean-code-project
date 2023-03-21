package encryption

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func CompareHashAndPassword(hash []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}

func FromHashToBase64(hash []byte) string {
	return base64.RawStdEncoding.EncodeToString(hash)
}

func FromBase64ToHash(plaintext string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(plaintext)
}
