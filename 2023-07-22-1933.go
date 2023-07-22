package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数字是否为偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用数组存储一组数字，并打印出来
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, n := range numbers {
		fmt.Println(n)
	}
}
