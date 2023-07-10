package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数相加的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 24
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用循环打印10个斐波那契数列的数字
	fibonacci := make([]int, 10)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < 10; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Println("Fibonacci:", fibonacci)

	// 练习5：求一个整数数组的和
	numbers := []int{1, 2, 3, 4, 5}
	sum = 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum of numbers:", sum)
}