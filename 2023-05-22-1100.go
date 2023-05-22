package main

import (
	"fmt"
)

func main() {
	// 练习1：打印1-10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：求1-100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：判断一个数是否为质数
	num := 17
	flag := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println(num, "是质数")
	} else {
		fmt.Println(num, "不是质数")
	}

	// 练习4：计算斐波那契数列第n项
	n := 10
	fib1 := 1
	fib2 := 1
	for i := 3; i <= n; i++ {
		temp := fib1 + fib2
		fib1 = fib2
		fib2 = temp
	}
	fmt.Println("斐波那契数列第", n, "项为", fib2)
}