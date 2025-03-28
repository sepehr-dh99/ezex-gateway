package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomCode(length uint8) string {
	const digits = "0123456789"
	code := make([]byte, length)

	for i := range code {
		num, err := rand.Int(rand.Reader, bigInt(len(digits)))
		if err != nil {
			// fallback to zero if crypto/rand fails (very rare)
			code[i] = '0'
			continue
		}
		code[i] = digits[num.Int64()]
	}

	return string(code)
}

func bigInt(n int) *big.Int {
	return big.NewInt(int64(n))
}
