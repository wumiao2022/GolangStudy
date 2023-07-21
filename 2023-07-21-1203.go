package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习4：使用函数交换两个变量的值
	a, b := 5, 10
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a = %d, b = %d\n", a, b)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}