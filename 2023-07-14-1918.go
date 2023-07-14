package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：输出1到10之间的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算1到100之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum of even numbers between 1 and 100:", sum)

	// 练习4：计算斐波那契数列的前20个数字
	num1, num2 := 0, 1
	fmt.Println(num1)
	fmt.Println(num2)
	for i := 2; i < 20; i++ {
		next := num1 + num2
		fmt.Println(next)
		num1, num2 = num2, next
	}
}