package main

import "fmt"

func main() {
	// 练习1: 声明一个整型变量，并打印其值
	var num int = 10
	fmt.Println(num)

	// 练习2: 定义一个字符串变量，并打印其值
	var str string = "Hello, Golang!"
	fmt.Println(str)

	// 练习3: 声明一个数组，并打印数组中的元素
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 练习4: 创建一个切片，并打印切片中的元素
	slice := []int{6, 7, 8, 9, 10}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// 练习5: 创建一个带有键值对的map，并打印map中的元素
	dict := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	for k, v := range dict {
		fmt.Println(k, v)
	}
}