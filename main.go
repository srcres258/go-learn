package main

import (
	"fmt"
	"reflect"
)

type User struct {}

func (User) Call(name string) {
	fmt.Println("我被调用了", name)
}

func Call(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()

	for i := 0; i < v.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println("方法名:", m.Name)
		if m.Name != "Call" {
			continue
		}
		method := v.Method(i)
		method.Call([]reflect.Value{
			reflect.ValueOf("zhangsan"),
		})
	}
}

func main() {
	s := User {}
	Call(&s)
}

