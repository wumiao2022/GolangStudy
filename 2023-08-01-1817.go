package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为奇数
	num := 5
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用循环计算1到100的和
	total := 0
	for i := 1; i <= 100; i++ {
		total += i
	}
	fmt.Println("Total:", total)

	// 练习5：使用函数交换两个变量的值
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
```