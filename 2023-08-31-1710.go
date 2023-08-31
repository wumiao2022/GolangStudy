package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1: 生成一个随机整数的切片，长度为10
	numbers := generateRandomNumbers(10)
	fmt.Println(numbers)

	// 练习2: 将切片中的奇数元素乘以2
	doubledNumbers := doubleOddNumbers(numbers)
	fmt.Println(doubledNumbers)

	// 练习3: 计算切片中元素的平均值
	average := calculateAverage(numbers)
	fmt.Println(average)
}

func generateRandomNumbers(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

func doubleOddNumbers(numbers []int) []int {
	doubledNumbers := make([]int, len(numbers))
	for i, num := range numbers {
		if num%2 != 0 {
			doubledNumbers[i] = num * 2
		} else {
			doubledNumbers[i] = num
		}
	}
	return doubledNumbers
}

func calculateAverage(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}