package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello, World!的代码
	fmt.Println("Hello, World!")


	// 练习2：求两个整数的和的代码
	var num1, num2 int
	num1 = 10
	num2 = 20
	sum := num1 + num2
	fmt.Println("The sum is:", sum)


	// 练习3：输出1到10所有偶数的代码
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}


	// 练习4：交换两个变量的值的代码
	a := 10
	b := 20
	temp := a
	a = b
	b = temp
	fmt.Println("a =", a)
	fmt.Println("b =", b)


	// 练习5：判断一个数是正数、负数还是0的代码
	num := 5
	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}