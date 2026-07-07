package auth

import (
	"crypto/rand"
	"math/big"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GenerateOTP(length int) (string, error) {
	b := make([]byte, length)

	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}

		b[i] = chars[n.Int64()]
	}

	return string(b), nil
}