package main

import "fmt"

type SingInterface interface {
	Sing()
	GetName() string
}

type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println(c.Name, "is singing")
}
func (c Chicken) GetName() string {
	return c.Name
}

type Cat struct {
	Name string
}

func (c Cat) Sing() {
	fmt.Println(c.Name, "is singing")
}
func (c Cat) GetName() string {
	return c.Name
}

func sing(c SingInterface) {
	c.Sing()
	fmt.Println(c.GetName())

	ch, ok := c.(Chicken)
	fmt.Println(ch, ok)

	switch server := c.(type) {
	case Chicken:
		fmt.Println("It's a chicken:", server.Name)
	case Cat:
		fmt.Println("It's a cat:", server.Name)
	default:
		fmt.Println("Unknown type")
	}
}

func Print(v any) {
	fmt.Println(v)
}

func main() {
	ch := Chicken{Name: "Clucky"}
	ca := Cat{Name: "Lihuamao"}
	sing(ch)
	sing(ca)
	Print(ch)
	Print(ca)
}

