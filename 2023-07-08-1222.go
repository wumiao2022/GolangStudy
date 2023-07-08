package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个整数的和、差、积和商
	num1 := 10
	num2 := 5
	sum := num1 + num2
	diff := num1 - num2
	product := num1 * num2
	quotient := num1 / num2

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %d\n", quotient)

	// 练习3: 判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	// 练习4: 求一个数的阶乘
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d is %d\n", n, factorial)

	// 练习5: 查找切片中的最大值和最小值
	numbers := []int{11, 4, 9, 7, 2, 15}
	max := numbers[0]
	min := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Printf("Max: %d\n", max)
	fmt.Printf("Min: %d\n", min)
}