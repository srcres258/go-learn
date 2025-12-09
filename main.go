package main

import (
	"reflect"
	"fmt"
)

func getType(obj any) {
	t := reflect.TypeOf(obj)
	switch t.Kind() {
	case reflect.String:
		fmt.Println("String")
	case reflect.Int:
		fmt.Println("Int")
	case reflect.Struct:
		fmt.Println("Struct")
	}
}

func main() {
	getType(114514)
	getType("114514")
	getType(struct {
		Name string
	}{})
}

