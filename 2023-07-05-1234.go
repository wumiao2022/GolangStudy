package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个整数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum =", sum)

	// 练习3：输出1到10的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习4：将字符串倒序输出
	str := "Golang"
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	fmt.Println(reverseStr)

	// 练习5：判断一个数是否为质数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println("Is", num, "prime?", isPrime)
}