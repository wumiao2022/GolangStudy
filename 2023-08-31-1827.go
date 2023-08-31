package main

import "fmt"

func main() {
	// 练习题1：声明并打印一个字符串变量
	var str string = "Hello, Golang!"
	fmt.Println(str)

	// 练习题2：声明并打印一个整数变量
	var num int = 42
	fmt.Println(num)

	// 练习题3：声明并打印一个浮点数变量
	var floatNum float64 = 3.14
	fmt.Println(floatNum)

	// 练习题4：声明并打印一个布尔变量
	var isTrue bool = true
	fmt.Println(isTrue)

	// 练习题5：声明并打印一个整数切片
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 练习题6：声明并打印一个键值对集合
	mapData := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(mapData)

	// 练习题7：声明一个函数并调用它
	sayHello()

	// 练习题8：调用内置的len函数并打印结果
	length := len(str)
	fmt.Println(length)
}

func sayHello() {
	fmt.Println("Hello, Golang!")
}
