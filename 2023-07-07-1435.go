package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 示例1：生成10个随机数，并打印出来
	fmt.Println("示例1:")
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}

	// 示例2：计算1到100之间所有数字的和
	fmt.Println("\n示例2:")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("总和为:", sum)

	// 示例3：打印1到100之间的所有奇数
	fmt.Println("\n示例3:")
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 示例4：使用冒泡排序算法对一个整数切片进行升序排序
	fmt.Println("\n示例4:")
	numbers := []int{4, 2, 1, 5, 3}
	fmt.Println("排序前:", numbers)

	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	fmt.Println("排序后:", numbers)
}