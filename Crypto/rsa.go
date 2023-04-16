package Crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
)

type r struct {}

func (r) Encode(args ...interface{}) map[string]interface{} {
	text, size := args[0].(string), len(args)

	encoding, bits := "base64", 4096

	if size == 2 {
		encoding = args[1].(string)
	} else if size == 3 {
		encoding, bits = args[1].(string), args[2].(int)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        panic(err)
    }

	publicKey := &privateKey.PublicKey

	ciphertext, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, publicKey, []byte(text), nil)
    if err != nil {
        panic(err)
    }

	return map[string]interface{} {
		"Text": Base.Encode(string(ciphertext), encoding),
		"privateKey": privateKey,
		"publicKey": publicKey,
		"encoding": encoding,
	}
}

func (r) Decode(data map[string]interface{}) string { // a single line nice 
	plaintext, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, data["privateKey"].(*rsa.PrivateKey), []byte(Base.Decode(data["Text"].(string), data["encoding"].(string))), nil)

	if err != nil {
		panic(err)
	}

	return string(plaintext)
}

var RSA = new(r)