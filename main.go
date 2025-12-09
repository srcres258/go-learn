package main

import (
	"fmt"
	// "encoding/json"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint |
		uint8 | uint16 | uint32 | uint64
}

func plus[T int | uint](n1, n2 T) T {
	return n1 + n2
}

type Response[T any] struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data T `json:"data"`
}

type User struct {
	Name string `json:"name"`
}

type UserInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	// plus(1, 2)
	// var u1, u2 = uint(1), uint(2)
	// plus(u1, u2)

	// user := Response {
	// 	Code: 0,
	// 	Msg: "成功",
	// 	Data: User {
	// 		Name: "张三",
	// 	},
	// }
	// byteData, _ := json.Marshal(user)
	// fmt.Println(string(byteData))
	//
	// userInfo := Response {
	// 	Code: 0,
	// 	Msg: "成功",
	// 	Data: UserInfo {
	// 		Name: "张三",
	// 		Age: 18,
	// 	},
	// }
	// byteData, _ = json.Marshal(userInfo)
	// fmt.Println(string(byteData))

	// var userResponse Response[User]
	// json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"张三"}}`), &userResponse)
	// fmt.Println(userResponse)
	// fmt.Println(userResponse.Data.Name)
	//
	// var userInfoResponse Response[UserInfo]
	// json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"张三","age":18}}`), &userInfoResponse)
	// fmt.Println(userInfoResponse)
	// fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age)

	// type MySlice [T int | string] []T
	// var mySlice = MySlice[int] {1,2,3}
	// fmt.Println(mySlice)

	type MyMap[T string, K any] map[T]K
	var myMap = MyMap[string, string]{
		"name": "12",
	}
	fmt.Println(myMap)
}

