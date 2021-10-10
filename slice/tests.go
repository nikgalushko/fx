package slice

import (
	"fmt"
	"reflect"
)


func equal[T any](actual, expected T, title string) {
	if !reflect.DeepEqual(actual, expected) {
		panic(fmt.Sprintf("actual %v != expected %v", actual, expected))
	} else {
		fmt.Println(title + " - OK")
	}
}
