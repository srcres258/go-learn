package main

import (
	"reflect"
	"fmt"
)

func getValue(obj any) {
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.String:
		fmt.Println("String", v.String())
	case reflect.Int:
		fmt.Println("Int", v.Int())
	case reflect.Struct:
		fmt.Println("Struct")
	}
}

func main() {
	getValue(114514)
	getValue("114514")
	getValue(struct {
		Name string
	}{})
}

