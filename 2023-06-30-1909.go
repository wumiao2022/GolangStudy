package main

import "fmt"

func main() {
	//练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	//练习2：交换两个变量的值
	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a, b)

	//练习3：计算两个数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	//练习4：判断一个数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	//练习5：打印1到10之间的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}