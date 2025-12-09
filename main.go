package main

import (
	"os"
	"bufio"
	"fmt"
	// "io"
)

func main() {
	// byteData, err := os.ReadFile("111")
	// byteData, err := os.ReadFile("/home/srcres/Coding/Learn/go-learn/test.txt")
	// byteData, err := os.ReadFile("test.txt")
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }
	// fmt.Println("File content:", string(byteData))

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// for {
	// 	var byteData = make([]byte, 15)
	// 	n, err := file.Read(byteData)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(string(byteData[:n]))
	// }

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// buf := bufio.NewReader(file)
	// for {
	// 	line, _, err := buf.ReadLine()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(string(line))
	// }

	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := bufio.NewScanner(file)
	buf.Split(bufio.ScanWords)
	var index int
	for buf.Scan() {
		index++
		fmt.Println(index, buf.Text())
	}
}

