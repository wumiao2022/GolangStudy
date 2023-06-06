package main

import (
	"fmt"
)

func main() {
	// 1. 连接两个字符串，并打印输出结果
	str1 := "Hello"
	str2 := "World"
	fmt.Println(str1 + str2)

	// 2. 判断一个数是否为偶数，并打印输出结果
	num1 := 10
	if num1%2 == 0 {
		fmt.Println(num1, "is even")
	} else {
		fmt.Println(num1, "is odd")
	}

	// 3. 创建一个长度为5的整型数组，并将数组中的元素都打印输出
	arr := [5]int{1, 2, 3, 4, 5}
	for _, num := range arr {
		fmt.Println(num)
	}

	// 4. 创建一个切片，并将切片的元素都打印输出
	slice := []string{"apple", "banana", "orange"}
	for _, item := range slice {
		fmt.Println(item)
	}

	// 5. 计算两个数的和，并打印输出结果
	num2 := 5
	num3 := 10
	fmt.Println(num2 + num3)
}