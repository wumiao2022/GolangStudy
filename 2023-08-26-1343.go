package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义并打印一个字符串变量
	message := "Welcome to Go programming!"
	fmt.Println(message)

	// 3. 定义一个整数变量，并进行四则运算
	num1 := 10
	num2 := 5
	sum := num1 + num2
	diff := num1 - num2
	mul := num1 * num2
	div := num1 / num2
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", diff)
	fmt.Printf("Product: %d\n", mul)
	fmt.Printf("Quotient: %d\n", div)

	// 4. 使用循环打印1到10之间的奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 5. 使用切片实现队列操作，包括入队和出队
	queue := []int{}
	queue = append(queue, 1) // 入队
	fmt.Println(queue)
	item := queue[0]        // 出队
	queue = queue[1:len(queue)]
	fmt.Println(item)
	fmt.Println(queue)
}
