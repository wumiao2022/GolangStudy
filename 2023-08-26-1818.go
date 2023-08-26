package main

import "fmt"

func main() {
	// 示例1：输出Hello World!
	fmt.Println("Hello, World!")

	// 示例2：整数相加
	a := 3
	b := 5
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	// 示例3：浮点数相乘
	x := 3.5
	y := 2.8
	product := x * y
	fmt.Printf("%.2f * %.2f = %.2f\n", x, y, product)

	// 示例4：字符串拼接
	str1 := "Hello"
	str2 := "World"
	concatenatedStr := str1 + ", " + str2 + "!"
	fmt.Println(concatenatedStr)

	// 示例5：条件判断
	num := 7
	if num > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is less than or equal to 10")
	}

	// 示例6：循环输出
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}