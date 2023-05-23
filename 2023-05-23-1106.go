package main

import (
	"fmt"
)

func main() {
	// 练习1: 输出10个斐波那契数列的值
	for i, a, b := 0, 0, 1; i < 10; i, a, b = i+1, b, a+b {
		fmt.Println(b)
	}

	// 练习2: 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3: 判断数字2到100之间有哪些素数并输出
	for i := 2; i <= 100; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
	
	// 练习4: 在命令行输入一个整数，然后逆序输出其数字
	var num, reverse int
	fmt.Print("请输入一个整数：")
	fmt.Scan(&num)
	for num != 0 {
		reverse = reverse*10 + num%10
		num = num / 10
	}
	fmt.Println(reverse)
}