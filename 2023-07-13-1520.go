package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个整数的和
	a := 5
	b := 3
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习4：交换两个变量的值
	x := 10
	y := 20
	fmt.Println("Before swap: x =", x, "y =", y)
	x, y = y, x
	fmt.Println("After swap: x =", x, "y =", y)

	// 练习5：使用循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}