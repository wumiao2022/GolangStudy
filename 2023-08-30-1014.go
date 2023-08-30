package main

import "fmt"

func main() {
	// 练习1: 输出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2: 输出两个数的和
	a := 10
	b := 5
	fmt.Println(a + b)

	// 练习3: 判断一个数是否为偶数，是则输出 "Even"，否则输出 "Odd"
	num := 7
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}