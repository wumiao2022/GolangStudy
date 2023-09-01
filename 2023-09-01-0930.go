package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum =", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 25
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：交换两个变量的值并打印结果
	a := 10
	b := 20
	fmt.Println("Before swap: a =", a, "b =", b)
	a, b = b, a
	fmt.Println("After swap: a =", a, "b =", b)

	// 练习5：遍历打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6：使用函数计算阶乘并打印结果
	n := 5
	factorial := calcFactorial(n)
	fmt.Println(n, "factorial is", factorial)
}

func calcFactorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * calcFactorial(n-1)
	}
}