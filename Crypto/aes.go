package Crypto

import (
	"io"
	"strings"
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"
)

type aes128 struct{}

func (aes128) Encode(args ...interface{}) map[string]string {
	text, key, size := []byte(args[0].(string)), []byte(MD5(args[1].(string))), len(args)

	Encoding := "base64"
	if size == 3 {
		Encoding = strings.ToLower(args[2].(string))
	} else if size < 2 || size > 3 {
		panic("aes-abrupt: AES.Encode invalid number of arguments")
	}

	c, err := aes.NewCipher(key)

	if err != nil {
        panic(err)
    }
	
	gcm, err := cipher.NewGCM(c)

    if err != nil {
        panic(err)
    }
	
    nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        panic(err)
    }
	
	out := string(gcm.Seal(nonce, nonce, text, nil))
	
    switch(Encoding) {
		case "base64":
			out = Base64.Encode(out)
		case "base32":
			out = Base32.Encode(out)
		case "hex":
			out = Hex.Encode(out)
	}

	return map[string]string{
		"Text": out,
		"Key": string(key),
		"Encoding": Encoding,
	}
}

func (aes128) Decode(args ...interface{}) string {
	size := len(args)
	var Encoding, Output, Key string
	Encoding = "base64"
	if size == 1 {
		data := args[0].(map[string]string)
		Output, Key, Encoding = data["Text"], data["Key"], data["Encoding"]
	} else if size == 2 || size == 3 {
		Output, Key = args[0].(string), args[1].(string)
		if size == 3 {
			Encoding = args[2].(string)
		}
	} else {
		panic("aes-abrupt: AES.Decode invalid number of arguments")
	}

	switch(Encoding) {
		case "base64":
			Output = Base64.Decode(Output)
		case "base32":
			Output = Base32.Decode(Output)
		case "hex":
			Output = Hex.Decode(Output)
	}
	
	text, key := []byte(Output), []byte(Key)
	
	c, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        panic(err)
    }

    nonceSize := gcm.NonceSize()
    if len(text) < nonceSize {
        panic(err)
    }

    nonce, text := text[:nonceSize], text[nonceSize:]
    out, err := gcm.Open(nil, nonce, text, nil)
    if err != nil {
        panic(err)
    }
	
	return string(out)
}

var AES = new(aes128)
