package String 

import (
	"fmt"
	"reflect"
)

func Reverse(String string) string {
	var Return string

	for i := len(String)-1; i >= 0; i-- {
		Return += string(String[i])
	}
	
	return Return
}

func Comma(arguments ...interface{}) interface{} {
	var String interface{} = arguments[0]

	switch reflect.TypeOf(String).Kind() {
		case reflect.String:
			String = String.(string)
		case reflect.Int:
			String = fmt.Sprint(String.(int))
		default:
			String = false
	}

	var ArgumentSize = len(arguments)
	var Prefix string
	if ArgumentSize == 2 {
		var Raw interface{} = arguments[1]
		if reflect.TypeOf(Raw).Kind() != reflect.String {
			return false
		}
		Prefix = Raw.(string)

	} else {
		if ArgumentSize != 1 {
			return false
		}
	}

	if String != false {
		var Assert string = String.(string)
		var Size int = len(Assert)-1
		var Return string
		for i := 0; i <= Size; i++ {
			var Letter string = string(Assert[i])
			if (i - (len(Assert)%3)) % 3 == 0 && i != 0 {
				Return += "," + Letter
			} else {
				Return += Letter
			}
		}
		return Prefix + Return
	}
	
	return String
}