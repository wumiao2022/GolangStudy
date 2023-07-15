package main

import "fmt"

func main() {
	// 练习1: 打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 变量交换
	a := 10
	b := 20
	fmt.Println("Before Swap - a:", a, "b:", b)
	a, b = b, a
	fmt.Println("After Swap - a:", a, "b:", b)

	// 练习3: 判断奇偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4: 循环计算阶乘
	factorial := 1
	num := 5
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", num, "is", factorial)

	// 练习5: 切片操作
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Original numbers:", numbers)
	fmt.Println("First 3 numbers:", numbers[:3])
	fmt.Println("Last 3 numbers:", numbers[3:])
	fmt.Println("Middle 2 numbers:", numbers[2:4])

	// 练习6: 函数返回多个值
	sum, prod := calculate(2, 3)
	fmt.Println("Sum:", sum)
	fmt.Println("Product:", prod)
}

// 函数返回两个数的和与乘积
func calculate(a, b int) (int, int) {
	return a + b, a * b
}