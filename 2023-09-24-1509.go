package main

import (
	"fmt"
	"math"
)

func main() {
	// 练习1：计算两个数的和并打印结果
	a := 5
	b := 3
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习2：计算平方根并打印结果
	x := 16.0
	sqrt := math.Sqrt(x)
	fmt.Println("Square root:", sqrt)

	// 练习3：将字符串反转并打印结果
	str := "Hello, World!"
	reversedStr := reverseString(str)
	fmt.Println("Reversed string:", reversedStr)
}

func reverseString(str string) string {
	// 将字符串转换为字节数组
	bytes := []byte(str)

	// 使用双指针法逆序交换字符
	left := 0
	right := len(str) - 1
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}

	// 将逆序后的字节数组转换为字符串并返回
	return string(bytes)
}
