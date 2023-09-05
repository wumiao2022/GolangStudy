package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. 判断一个数是奇数还是偶数
	number := 15
	if number%2 == 0 {
		fmt.Println(number, "is even")
	} else {
		fmt.Println(number, "is odd")
	}

	// 4. 输出1到10的平方数
	for i := 1; i <= 10; i++ {
		square := i * i
		fmt.Println(square)
	}

	// 5. 求一个数组中所有元素的和
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println("Sum of array:", sum)
}