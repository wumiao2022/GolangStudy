package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")
	
	// 练习2：计算两个数的和并输出
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("sum =", sum)
	
	// 练习3：判断一个数是否为偶数并输出结果
	num := 7
	if num % 2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	
	// 练习4：循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}