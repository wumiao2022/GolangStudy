package main

import "fmt"

func main() {
	// 练习1：打印出1到10的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}

	// 练习2：将两个数相加并输出结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println(sum)

	// 练习3：判断一个数是否是素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

	// 练习4：将一个字符串反转输出
	str := "Hello, World!"
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	fmt.Println(reverseStr)
}