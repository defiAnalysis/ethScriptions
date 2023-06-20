package utils

import (
	"encoding/hex"
	"strings"
)

func HexToString(hexString string) (string, error) {
	if strings.HasPrefix(hexString, "0x") {
		hexString = hexString[2:]
	}

	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	str := string(decoded)

	return str, nil
}
