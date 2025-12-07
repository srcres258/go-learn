package main

import "fmt"

type Class struct {
	Name string
}

type Student struct {
	Class
	Name string
}

func (s *Student)Study() {
	fmt.Printf("%s is learning...\n", s.Name)
}
func (s *Student)Info() {
	fmt.Printf("Name: %s, Class: %s\n", s.Name, s.Class.Name);
}

func (s *Student) SetName(name string) {
	s.Name = name
}

func main() {
	c1 := Class {
		Name: "Class 1",
	}

	s1 := Student { Class: c1, Name: "Alice" }
	s1.Study()
	s1.Info()

	s2 := Student { Class: c1, Name: "Bob" }
	s2.Study()
	s2.Info()

	fmt.Printf("%p\n", &s2)
	s2.SetName("Charlie")
	s2.Info()
}

