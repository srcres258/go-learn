package main

import (
	"os"
	// "bufio"
	"fmt"
	// "io"
)

func main() {
	// file, err := os.OpenFile("w.txt", os.O_CREATE | os.O_WRONLY, 0777)
	// file, err := os.OpenFile("w.txt", os.O_CREATE | os.O_RDWR, 0777)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// // file.Write([]byte("你好"))
	// byteData, err := io.ReadAll(file)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(byteData))

	// err := os.WriteFile("w1.txt", []byte("hello world"), 0666)
	// fmt.Println(err)

	// rFile, err := os.Open("/home/srcres/Pictures/0303f50b6eabbbce23_n.jpg")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer rFile.Close()
	// wFile, err := os.OpenFile("pic.jpg", os.O_CREATE | os.O_WRONLY, 0777)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// io.Copy(wFile, rFile)

	dir, err := os.ReadDir("/home/srcres/Coding/Learn/go-learn")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, entry := range dir {
		info, _ := entry.Info()
		fmt.Println(entry.IsDir(), entry.Name(), info.Name(), info.Size(), info.Mode())
	}
}

