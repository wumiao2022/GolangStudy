package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：判断一个数是偶数还是奇数
	num := 7
	if num % 2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习3：计算两个数的和、差、积、商
	a := 8
	b := 3
	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)

	// 练习4：遍历打印10以内的所有整数
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 练习5：计算1到10之间所有整数的平均值
	sum = 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	average := float64(sum) / 10
	fmt.Printf("Average: %.2f\n", average)
}