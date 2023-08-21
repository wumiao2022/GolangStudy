package main

import "fmt"

func main() {
	// 练习1: 打印Hello World
	fmt.Println("Hello World")

	// 练习2: 定义一个整数变量并打印出来
	num := 10
	fmt.Println(num)

	// 练习3: 使用循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4: 定义一个切片并打印出来
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 练习5: 使用函数交换两个整数的值并打印结果
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
