package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")
	
	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)
	
	// 练习3: 判断一个数是奇数还是偶数
	num := 5
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习4: 判断一个年份是否是闰年
	year := 2021
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a Leap Year")
	}

	// 练习5: 计算一个数的阶乘
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial:", factorial)

	// 练习6: 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}