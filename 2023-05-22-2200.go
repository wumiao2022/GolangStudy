package main

import (
	"fmt"
)

func main() {
	// 练习1：输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：打印1-10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算1+2+...+10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习4：判断一个数是否为素数
	num := 13
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d是素数\n", num)
	} else {
		fmt.Printf("%d不是素数\n", num)
	}

	// 练习5：翻转一个字符串
	str := "Hello, World!"
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	fmt.Println(reverseStr)
}