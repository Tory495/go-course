package generator

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

const DefaultBytesLength = 32

func GenerateSessionId() (string, error) {
	bytes := make([]byte, DefaultBytesLength)
	_, err := rand.Read(bytes)

	if err != nil {
		return "", fmt.Errorf("can't generate session id: %w", err)
	}

	return hex.EncodeToString(bytes), nil
}

const DefaultSmsValue = 1000000

func GenerateSmsCode() (int, error) {
	max := big.NewInt(DefaultSmsValue)
	n, err := rand.Int(rand.Reader, max)

	if err != nil {
		return 0, fmt.Errorf("cannot generate sms code: %w", err)
	}

	return int(n.Int64()), nil
}
