package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 练习4：使用循环计算1到10的累加和并打印结果
	total := 0
	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Println("Sum of 1 to 10:", total)

	// 练习5：使用数组存储和打印出多个姓名
	names := []string{"Alice", "Bob", "Charlie", "Daisy"}
	for _, name := range names {
		fmt.Println(name)
	}
}