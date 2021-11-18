package Crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/base32"
	"encoding/hex"
)


func digest(encode string, data []byte) string {
	var result string
	switch encode {
		case "hex":
			result = hex.EncodeToString(data)
		case "base64":
			result = base64.StdEncoding.EncodeToString(data)
		case "base32":
			result = base32.StdEncoding.EncodeToString(data)
	}
	return result
}

func construct(arguments []string) (string, string) {
	var text, encode string = arguments[0], "hex"

	if len(arguments) == 2 {
		encode = arguments[1]
	}

	return text, encode
}

func SHA256(arguments ...string) string {
	var text, encode string = construct(arguments)
	hash := sha256.New()
	hash.Write([]byte(text))
	return digest(encode, hash.Sum(nil))
}

func SHA512(arguments ...string) string {
	var text, encode string = construct(arguments)
	hash := sha512.New()
	hash.Write([]byte(text))
	return digest(encode, hash.Sum(nil))
}

func MD5(arguments ...string) string {
	var text, encode string = construct(arguments)
	hash := md5.Sum([]byte(text))
	return digest(encode, hash[:])
}
