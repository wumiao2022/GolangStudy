package main

import "fmt"

func main() {
	// 练习1：输出Hello, world!
	fmt.Println("Hello, world!")

	// 练习2：变量交换
	a := 1
	b := 2
	a, b = b, a
	fmt.Println("a =", a) // 2
	fmt.Println("b =", b) // 1

	// 练习3：循环打印数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：字符串反转
	str := "Hello, world!"
	reverseStr := ""
	for _, char := range str {
		reverseStr = string(char) + reverseStr
	}
	fmt.Println(reverseStr)

	// 练习5：求两个数的最大公约数
	num1 := 20
	num2 := 30
	for num2 != 0 {
		num1, num2 = num2, num1%num2
	}
	fmt.Println("最大公约数:", num1)
}