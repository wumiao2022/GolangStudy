package main

import "fmt"

func main() {
	// 实例1：打印Hello, World!
	fmt.Println("Hello, World!")
	
	// 实例2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)
	
	// 实例3：判断一个数是否为偶数
	num := 24
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	
	// 实例4：交换两个变量的值
	a := 5
	b := 10
	fmt.Println("Before swapping, a =", a, "and b =", b)
	a, b = b, a
	fmt.Println("After swapping, a =", a, "and b =", b)
}