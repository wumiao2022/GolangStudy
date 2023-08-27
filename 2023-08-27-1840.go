package main

import "fmt"

func main() {
	// 练习1：输出Hello World！
	fmt.Println("Hello World!")

	// 练习2：打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习4：判断一个数是否是偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习5：求一个数的阶乘
	factorial := 1
	n := 5
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println(factorial)
}