package main

import "fmt"

func main() {
	// 练习1：使用 fmt 包输出 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：定义一个变量 univ，并赋值为 "Golang"
	univ := "Golang"
	fmt.Println(univ)

	// 练习3：定义两个整型变量 a 和 b，并分别赋值为 10 和 5，然后输出 a 和 b 的和、差、积和商
	a := 10
	b := 5
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
}