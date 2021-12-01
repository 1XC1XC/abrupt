package Crypto

import (
	"strings"
	"unicode"
	// "fmt"
)
// Total:!"&'()+,-./0123456789:=?@abcdefghijklmnopqrstuvwxyz

var MC = map[string]string{"!":"-.-.--","\"":".-..-.","&":".-...","'":".----.","(":"-.--.",")":"-.--.-","+":".-.-.",",":"--..--","-":"-....-",".":".-.-.-","/":"-..-.","0":"-----","1":".----","2":"..---","3":"...--","4":"....-","5":".....","6":"-....","7":"--...","8":"---..","9":"----.",":":"---...","=":"-...-","?":"..--..","@":".--.-.","a":".-","b":"-...","c":"-.-.","d":"-..","e":".","f":"..-.","g":"--.","h":"....","i":"..","j":".---","k":"-.-","l":".-..","m":"--","n":"-.","o":"---","p":".--.","q":"--.-","r":".-.","s":"...","t":"-","u":"..-","v":"...-","w":".--","x":"-..-","y":"-.--","z":"--.."}

var MCR = func() map[string]string {
	r := make(map[string]string)

	for i, v := range MC {
	    r[v] = i
	}
	return r 
}()

type M struct {}

func (_ M) Encode(text string) string {
	var out string
	for _, v := range text {
		out += MC[string(unicode.ToLower(v))] + " "
	}
	return strings.TrimSuffix(out, " ")
}

func (_ M) Decode(text string) string {
	var out string
	for _, v := range strings.Split(text, " ") {
		if v == "" {
			out += " "
		} else {
			out += MCR[v]
		}
	}
	return out 
}

var Morse = new(M)
