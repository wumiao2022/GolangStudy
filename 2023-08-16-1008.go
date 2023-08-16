package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个整数是否为偶数
	num := 8
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 练习4：使用循环计算1到100的和
	sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of numbers from 1 to 100:", sum)

	// 练习5：将字符串数组反转
	strArr := []string{"Hello", "World", "Golang"}
	reversedArr := make([]string, len(strArr))
	for i := 0; i < len(strArr); i++ {
		reversedArr[i] = strArr[len(strArr)-1-i]
	}
	fmt.Println("Reversed array:", reversedArr)
}

输出结果:
Hello, World!
Sum: 15
8 is even.
Sum of numbers from 1 to 100: 5050
Reversed array: [Golang World Hello]