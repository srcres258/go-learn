package main

import "fmt"

type Animal struct {
	Name string
}

type Animal1 struct {
	Name1 string
}

type Cat struct {
	Animal
	Animal1
}

func main() {
	var animal = Animal { Name: "Cat" }
	cat := Cat {
		Animal: animal,
	}
	fmt.Println(cat.Name, cat.Name1)
}

