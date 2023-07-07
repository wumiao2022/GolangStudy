package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：遍历数组并打印元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习5：求一个数的阶乘
	number := 5
	factorial := 1
	for i := 1; i <= number; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", number, "is", factorial)
}
