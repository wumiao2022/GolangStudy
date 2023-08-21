package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Printf("Sum: %d\n", sum)

	// 练习3：判断一个数是否为偶数
	number := 7
	if number%2 == 0 {
		fmt.Printf("%d is even\n", number)
	} else {
		fmt.Printf("%d is odd\n", number)
	}

	// 练习4：将一个字符串反转并输出
	str := "Hello, Golang"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)

	// 练习5：打印1到100之间的所有素数
	for num := 2; num <= 100; num++ {
		isPrime := true
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println()
}
