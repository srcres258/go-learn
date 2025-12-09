package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name1 string `big:"-"`
	Name2 string
}

func SetStruct(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		big := t.Field(i).Tag.Get("big")
		if big == "" {
			continue
		}
		value.SetString(value.String() + " modified")
	}
}

func main() {
	s := User { Name1: "name1", Name2: "name2" }
	SetStruct(&s)
	fmt.Println(s)
}

