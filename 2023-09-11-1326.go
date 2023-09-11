package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量交换
	a := 5
	b := 10

	// 使用中间变量来交换a和b的值
	temp := a
	a = b
	b = temp

	fmt.Println("a =", a) // 预期输出: a = 10
	fmt.Println("b =", b) // 预期输出: b = 5

	// 练习3：函数差值
	num1 := 10
	num2 := 5

	diff := difference(num1, num2)
	fmt.Println("diff =", diff) // 预期输出: diff = 5

	// 练习4：字符串反转
	str := "Hello, World!"

	reversedStr := reverseString(str)
	fmt.Println("reversedStr =", reversedStr) // 预期输出: reversedStr = !dlroW ,olleH
}

// difference 函数返回两个数之间的差值
func difference(a, b int) int {
	return a - b
}

// reverseString 函数返回字符串的反转结果
func reverseString(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```