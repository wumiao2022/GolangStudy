package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并输出结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：计算圆的面积并输出结果
	radius := 5.0
	area := 3.14 * radius * radius
	fmt.Println("Area of the circle:", area)

	// 练习5：将一个字符串反转并输出结果
	str := "Hello, Go"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed String:", reversedStr)
}

// 请自行运行以上代码来进行每日练习，希望对你的学习有所帮助！