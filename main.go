package main

import "fmt"

type UserInfo struct {
	Name string `json:"name"`
}

func (this *UserInfo)SetName(name string) {
	this.Name = name
}

func main() {
	user := UserInfo{
		Name: "Alice",
	}
	user.SetName("Bob")
	fmt.Println(user.Name)
}

