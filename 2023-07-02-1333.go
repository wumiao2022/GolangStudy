package main

import "fmt"

func main() {
	// 1. 打印 Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算两个整数的和
	num1, num2 := 10, 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. 使用 for 循环打印 1 到 10 的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 4. 定义一个切片并遍历输出其中的元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Print(num, " ")
	}
	fmt.Println()

	// 5. 判断一个数是否为偶数，并输出结果
	num := 8
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
