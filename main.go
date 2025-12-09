package main

import (
	"reflect"
	"fmt"
)

func setValue(obj any, value any) {
	v1 := reflect.ValueOf(obj)
	v2 := reflect.ValueOf(value)
	if v1.Elem().Kind() != v2.Kind() {
		return
	}
	switch v1.Elem().Kind() {
	case reflect.String:
		v1.Elem().SetString(value.(string))
	case reflect.Int:
		v1.Elem().SetInt(v2.Int())
	}
}

func main() {
	var name = "田所"
	var age = 24
	setValue(&name, "德川")
	setValue(&age, 25)
	fmt.Println(name, age)
}

