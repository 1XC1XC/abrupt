package Base64


import (
	"encoding/base64"
)


func Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Decode(text string) string {
	decoded, _ := base64.StdEncoding.DecodeString(text)
	return string(decoded)
}