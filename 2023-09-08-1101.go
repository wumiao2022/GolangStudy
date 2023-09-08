package main

import "fmt"

func main() {
	// 练习1：声明并打印一个整数变量
	var num int = 10
	fmt.Println(num)

	// 练习2：声明并打印一个字符串变量
	var str string = "Hello, World!"
	fmt.Println(str)

	// 练习3：声明一个整数数组，并打印数组中的每个元素
	var arr = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 练习4：声明一个slice，并打印slice中的每个元素
	slice := []string{"apple", "banana", "orange"}
	for _, item := range slice {
		fmt.Println(item)
	}

	// 练习5：声明一个map，并打印map中的每个键值对
	mymap := map[string]int{"apple": 1, "banana": 2, "orange": 3}
	for key, value := range mymap {
		fmt.Println(key, value)
	}
}
