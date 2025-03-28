package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode"
)

func TestGenerateRandomCode(t *testing.T) {
	t.Run("GenerateRandomCode", func(t *testing.T) {
		code := GenerateRandomCode(6)

		assert.Equal(t, 6, len(code), "code length should be 6")

		for _, ch := range code {
			assert.True(t, unicode.IsDigit(ch), "each character should be a digit")
		}
	})

	t.Run("GenerateRandomCodeWithLength", func(t *testing.T) {
		for i := uint8(1); i <= 10; i++ {
			code := GenerateRandomCode(i)
			assert.Equal(t, int(i), len(code), "code length should match input")
		}
	})

	t.Run("GenerateRandomCodeWithUniqueness", func(t *testing.T) {
		code1 := GenerateRandomCode(6)
		code2 := GenerateRandomCode(6)

		// It's possible but very rare for them to match
		assert.NotEqual(t, code1, code2, "codes should be different (usually)")
	})

}
