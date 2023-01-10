package lib

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	ret := make([]byte, n)

	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}

		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func MustParseUint64(s *string) uint64 {
	var res uint64

	if s != nil {
		if parsedUint, err := strconv.ParseUint(*s, 0, 0); err == nil {
			res = parsedUint
		}
	}

	return res
}
