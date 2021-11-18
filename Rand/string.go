package Rand 

import "strings"

func Bool() bool {
	return Int(1)==1
}

func Str(arguments ...interface{}) string {
	Length, Letters := 5, false

	for i,v := range arguments {
		switch i {
			case 0:
				Length = v.(int)
			case 1:
				Letters = v.(bool)
		}
	}
	
	var Range []int

	if Letters {
		Range = []int{97,122}
	} else {
		Range = []int{33,126}
	}

	var Result string
	for i := 0; i < Length; i++ {
		Letter := string(Int(Range[0], Range[1])) 
		
		if Letters && Bool() {
			Letter = strings.ToUpper(Letter)
		}

		Result += Letter
	}
	
	return Result
}