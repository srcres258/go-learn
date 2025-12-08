package main

import (
	"fmt"
	"errors"
)

func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a / b
	return
}

func server() (res int, err error) {
	res, err = div(2, 0)
	if err != nil {
		// 把错误向上传递
		return
	}

	// 执行其他的逻辑
	res++
	res += 2

	return
}

func main() {
	res, err := server()
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	fmt.Println("结果:", res)
}

