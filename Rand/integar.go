/*
	I don't know if this is a poor usage I believe there probably is
	a more reliable method to writing this section. ( I'm bad at compiled languages )
*/

package Rand

import (
	"math/rand"
	"fmt"
	"time"
	"reflect"
)


func init() {
	rand.Seed(time.Now().UnixNano())
}

func construct(arguments []int) (int, int) {
	min, max := 1,5

	switch len(arguments) { 
		case 1:
			min = 0
			max = arguments[0]
		case 2:
			min = arguments[0]
			max = arguments[1]
	}

	if (max > min) {
		return min, max
	} else {
        fmt.Println("rand.Int(): Minimum > Maximum")
		return 1,5
	}
}

func Int(arguments ...int) int {
	min, max := construct(arguments)
	return rand.Intn(max - min + 1) + min
}

func Float(arguments ...int) float64 {
	min, max := construct(arguments)
	fmin, fmax := float64(min), float64(max)
	return (rand.Float64() * (fmax - fmin)) + fmin
}

func Array(argument interface{}) interface{} { // Interfacing is cool
	Array := argument
	Kind := reflect.TypeOf(Array).Kind()

	if !(Kind == reflect.Array || Kind == reflect.Slice) {
		fmt.Println("rand.Array: Array Type Required")
		return false
	}
	
	Reflection := reflect.ValueOf(Array)
	Size := Reflection.Len()
	
	if Size == 1 {
		return Reflection.Index(0)
	} else if Size == 0 {
		fmt.Println("rand.Array: Empty")
		return false
	}

	return Reflection.Index(Int(0,Size-1))
}