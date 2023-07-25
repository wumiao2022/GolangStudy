package main

import "fmt"

func main() {
	// 1. 计算两个整数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println(sum)

	// 2. 将一个字符串反转并输出结果
	str := "Hello, world!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println(reversedStr)

	// 3. 判断一个整数是否为偶数，并输出结果
	num := 15
	isEven := num%2 == 0
	fmt.Println(isEven)

	// 4. 将一个整数转换为二进制表示并输出结果
	num = 26
	binary := ""
	for num > 0 {
		remainder := num % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		num = num / 2
	}
	fmt.Println(binary)
}