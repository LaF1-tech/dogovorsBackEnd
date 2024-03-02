package randomizer

import (
	"crypto/rand"
	"math/big"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomString(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}

	b := make([]rune, length)
	for i := range b {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			return "", err
		}

		b[i] = letterRunes[int(nBig.Int64())]
	}

	return string(b), nil
}
