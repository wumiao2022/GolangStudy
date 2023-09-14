package main

import "fmt"

func main() {
	// 1. 输出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 计算并输出 2 + 3 的结果
	num1 := 2
	num2 := 3
	result := num1 + num2
	fmt.Println(result)

	// 3. 定义一个字符串变量，并输出该字符串
	message := "Hello, Golang!"
	fmt.Println(message)

	// 4. 定义一个整数变量，并输出该整数
	number := 10
	fmt.Println(number)

	// 5. 使用循环打印出 1 到 10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 6. 使用条件语句判断一个数字是否大于 5，如果是则输出 "大于 5"，否则输出 "小于等于 5"
	num := 8
	if num > 5 {
		fmt.Println("大于 5")
	} else {
		fmt.Println("小于等于 5")
	}

	// 7. 定义一个切片，并输出该切片的长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("长度:", len(slice))
	fmt.Println("容量:", cap(slice))
}