package Base32


import (
	"encoding/base32"
)


func Encode(text string) string {
	return base32.StdEncoding.EncodeToString([]byte(text))
}

func Decode(text string) string {
	decoded, _ := base32.StdEncoding.DecodeString(text)
	return string(decoded)
}