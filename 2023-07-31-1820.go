package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：将两个整数相加并输出结果
	num1 := 10
	num2 := 25
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 12
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用for循环输出1到10的平方数
	for i := 1; i <= 10; i++ {
		fmt.Println(i, i*i)
	}

	// 练习5：调用一个函数并传递参数
	name := "Alice"
	printName(name)

	// 练习6：使用数组存储一组整数，并输出数组中的元素
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
}

// 定义一个函数用于输出姓名
func printName(name string) {
	fmt.Println("Name:", name)
}
