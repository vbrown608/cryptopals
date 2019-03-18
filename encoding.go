package main

import (
	"encoding/base64"
)

func HexTo64(hex string) string {
	bytes := HexToBytes(hex)
	return base64.StdEncoding.EncodeToString(bytes)
}

func HexToBytes(hex string) []byte {
	result := make([]byte, len(hex)/2)
	for i := 0; i < len(hex); i += 2 {
		high := hexByte(hex[i]) << 4
		low := hexByte(hex[i+1])
		result[i/2] = high | low
	}
	return result
}

func hexByte(c byte) byte {
	if c >= 48 && c < 58 {
		return c - 48
	}
	if c >= 97 && c < 103 {
		return c - 87
	}
	panic("Hex character out of range")
}
