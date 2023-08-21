package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("Sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3：判断一个数是否为偶数
	num := 23
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用条件循环打印1到10的偶数
	i := 1
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i)
		}
		i++
	}
}