package main

import "fmt"

func main() {
	// 1. 计算两个数的和并输出结果
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 2. 判断一个数是否为偶数，并输出结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 3. 使用循环打印出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 4. 定义一个函数，输入两个参数并返回它们的乘积
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Product:", multiply(4, 5))

	// 5. 创建一个包含5个元素的字符串数组，并遍历输出每个元素
	strArray := [5]string{"apple", "banana", "orange", "grape", "watermelon"}
	for _, str := range strArray {
		fmt.Println(str)
	}
}