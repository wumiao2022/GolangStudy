package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：判断一个数是奇数还是偶数
	num := 7

	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习4：计算1到100的和
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("1到100的和:", sum)

	// 练习5：求两个数的最大值
	a, b := 12, 8

	if a > b {
		fmt.Println("最大值:", a)
	} else {
		fmt.Println("最大值:", b)
	}
}