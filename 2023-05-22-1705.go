package main

import (
	"fmt"
)

func main() {
	// 练习1：求1到100所有数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100所有数的和为:", sum)

	// 练习2：判断一个数是否为偶数
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习3：计算一个数的阶乘
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println(n, "的阶乘为:", factorial)
}