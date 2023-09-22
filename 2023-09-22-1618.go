package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1: 生成一个随机整数切片，并打印出来
	numbers := make([]int, 10)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
	}
	fmt.Println(numbers)

	// 练习2: 计算切片中的所有元素之和
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// 练习3: 判断切片中是否有大于50的元素
	hasLargeNum := false
	for _, num := range numbers {
		if num > 50 {
			hasLargeNum = true
			break
		}
	}
	fmt.Println("Has large number:", hasLargeNum)

	// 练习4: 将切片中的元素按照从大到小的顺序进行排序
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] < numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	fmt.Println(numbers)
}