package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：求两个整数的和
	num1 := 5
	num2 := 10
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习3：计算圆的面积
	radius := 5.0
	area := 3.14 * radius * radius
	fmt.Println("The area of the circle is:", area)

	// 练习4：判断一个数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is not an even number")
	}

	// 练习5：计算斐波那契数列的前n项
	n := 10
	first := 0
	second := 1
	fmt.Print("The Fibonacci Series is:", first, " ", second)
	for i := 3; i <= n; i++ {
		next := first + second
		fmt.Print(" ", next)
		first = second
		second = next
	}
	fmt.Println()
}
```

这是一个简单的练习代码，包括打印Hello World、求两个整数的和、计算圆的面积、判断一个数是否为偶数以及计算斐波那契数列的前n项。每个练习都有相应的输出结果。你可以运行这段代码来查看结果。