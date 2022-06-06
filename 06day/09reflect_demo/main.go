package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

func reflactType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
	fmt.Printf("type name: %v, type kind: %v\n", v.Name(), v.Kind())
}

func main() {
	var a float32 = 3.14
	reflactType(a)

	var b int64 = 100
	reflactType(b)

	var c = Cat{}
	reflactType(c)
}
