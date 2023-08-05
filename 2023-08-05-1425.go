package main

import "fmt"

func main() {
	// 练习一：输出Hello World
	fmt.Println("Hello World")

	// 练习二：计算两个数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println(sum)

	// 练习三：判断一个数是奇数还是偶数并输出结果
	num := 25
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习四：打印1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习五：输出一个九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}