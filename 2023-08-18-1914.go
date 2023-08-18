package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)

	// 练习3：判断一个数是偶数还是奇数
	num := 15
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	// 练习4：反转字符串
	str := "Hello, Golang"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Printf("Reversed string: %s\n", reversedStr)
}