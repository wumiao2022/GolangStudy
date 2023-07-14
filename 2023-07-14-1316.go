package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数数组
	numbers := generateRandomNumbers(10)
	
	// 打印数组元素
	fmt.Println(numbers)

	// 计算数组中的最大值和最小值
	max, min := findMaxAndMin(numbers)
	fmt.Printf("最大值：%d，最小值：%d\n", max, min)
	
	// 对数组进行排序
	sortedNumbers := sort(numbers)
	
	// 打印排序后的数组
	fmt.Println(sortedNumbers)
}

// 生成一个指定长度的随机数数组
func generateRandomNumbers(length int) []int {
	numbers := make([]int, length)
	for i := 0; i < length; i++ {
		numbers[i] = rand.Intn(100) // 随机生成0-100之间的整数
	}
	return numbers
}

// 查找数组中的最大值和最小值
func findMaxAndMin(numbers []int) (int, int) {
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
	return max, min
}

// 对数组进行排序
func sort(numbers []int) []int {
	length := len(numbers)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}