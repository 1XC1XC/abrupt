package Hex


import (
	"encoding/hex"
)

func Encode(text string) string {
	return hex.EncodeToString([]byte(text))
}

func Decode(text string) string {
	encode, _ := hex.DecodeString(text)
	return string(encode)
}