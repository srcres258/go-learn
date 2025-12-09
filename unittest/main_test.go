package main

import "fmt"
import "testing"
import "os"

func setup() {
	fmt.Println("Setting up tests...")
}
func teardown() {
	fmt.Println("Tearing down tests...")
}

func TestAdd2(t *testing.T) {
	fmt.Println("Running TestAdd2")
	t.Errorf("fail")
}

func TestMain(m *testing.M) {
	fmt.Println("Starting tests...")
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

