package main

import "fmt"

func main() {
	// 实例1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 实例2: 计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 实例3: 判断一个数是否为偶数
	num := 12
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 实例4: 求一个数的阶乘
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", n, "is", factorial)

	// 实例5: 反转字符串
	str := "Hello, Golang!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed string:", reversedStr)
}

注意：以上示例代码仅供练习参考，不包含任何解释或说明。如有疑问，请提问。