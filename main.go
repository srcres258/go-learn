package main

import "fmt"
import "encoding/json"

type User struct {
	Name string `json:"name"`
	Age int `json:"age,omitempty"`
	Password string `json:"-"`
}

func main() {
	user := User { Name: "Alice", Age: 0, Password: "123456" }
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
}

