package Crypto

import (
	"encoding/base64"
	"encoding/base32"
	"encoding/hex"
)

type b64 struct {}
type b32 struct {}
type b16 struct {}

func (b64) Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func (b64) Decode(text string) string {
	decoded, _ := base64.StdEncoding.DecodeString(text)
	return string(decoded)
}

func (b32) Encode(text string) string {
	return base32.StdEncoding.EncodeToString([]byte(text))
}

func (b32) Decode(text string) string {
	decoded, _ := base32.StdEncoding.DecodeString(text)
	return string(decoded)
}

func (b16) Encode(text string) string {
	return hex.EncodeToString([]byte(text))
}

func (b16) Decode(text string) string {
	encode, _ := hex.DecodeString(text)
	return string(encode)
}

var Base64 = new(b64)
var Base32 = new(b32)
var Hex = new(b16)
