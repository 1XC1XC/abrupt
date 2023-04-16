package Crypto

import (
	"encoding/base64"
	"encoding/base32"
	"encoding/hex"
)

type b struct {}
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

func (b) Encode(text string, encoding string) string { // for rsa
	switch(encoding) {
		case "base64":
			text = Base64.Encode(text)
		case "base32":
			text = Base32.Encode(text)
		case "hex":
			text = Hex.Encode(text)
	}

	return text
}

func (b) Decode(text string, encoding string) string {
	switch(encoding) {
		case "base64":
			text = Base64.Decode(text)
		case "base32":
			text = Base32.Decode(text)
		case "hex":
			text = Hex.Decode(text)
	}

	return text
}

var Base64 = new(b64)
var Base32 = new(b32)
var Hex = new(b16)
var Base = new(b)