package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("Sum is %d\n", sum)

	// 练习3：判断一个数是否为偶数
	number := 15
	if number%2 == 0 {
		fmt.Printf("%d is even\n", number)
	} else {
		fmt.Printf("%d is odd\n", number)
	}

	// 练习4：循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：计算斐波那契数列的前n项（0, 1, 1, 2, 3, 5, ...）
	n := 10
	first := 0
	second := 1
	fmt.Print("Fibonacci series: ", first, ", ", second)
	for i := 2; i < n; i++ {
		next := first + second
		fmt.Print(", ", next)
		first = second
		second = next
	}
	fmt.Println()
}