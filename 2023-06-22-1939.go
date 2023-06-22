package main

import (
	"fmt"
)

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：输出1-10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：输出斐波那契数列前10项
	fibonacci := []int{1, 1}
	for i := 2; i < 10; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println(fibonacci)

	// 练习4：在切片中查找某个元素的位置
	fruits := []string{"apple", "banana", "orange", "grape", "watermelon"}
	target := "orange"
	for i, fruit := range fruits {
		if fruit == target {
			fmt.Printf("%s found at index %d in the slice\n", target, i)
			break
		}
	}

	// 练习5：使用goroutine和channel实现生产者消费者模式
	c := make(chan int)

	// 生产者
	go func() {
		for i := 0; i < 10; i++ {
			c <- i // 往通道中写数据
		}
		close(c) // 关闭通道
	}()

	// 消费者
	for v := range c {
		fmt.Println(v) // 从通道中读数据
	}
}