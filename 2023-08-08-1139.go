package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算10的阶乘
	result := 1
	for i := 1; i <= 10; i++ {
		result *= i
	}
	fmt.Println("10的阶乘是:", result)

	// 练习3: 判断一个数是否为素数
	num := 7
	isPrime := true
	for i := 2; i <= num/2; i++ {
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

	// 练习4: 字符串反转
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("反转后的字符串:", reversedStr)
}