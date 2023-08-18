package main

import "fmt"

func main() {
	// 实现一个函数，交换两个变量的值，并打印交换前后的值
	a := 10
	b := 20
	fmt.Printf("交换前：a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("交换后：a=%d, b=%d\n", a, b)
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}