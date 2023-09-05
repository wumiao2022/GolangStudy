package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：生成10个0到100之间的随机数并计算它们的平均值
	var sum int
	for i := 0; i < 10; i++ {
		num := rand.Intn(101)
		sum += num
	}
	avg := float64(sum) / 10
	fmt.Printf("平均值：%f\n", avg)

	// 练习2：使用切片实现冒泡排序
	nums := []int{5, 2, 8, 3, 9, 1}
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println("排序后：", nums)

	// 练习3：编写一个函数，接收一个整数切片，并返回切片中的最大值和最小值
	arr := []int{7, 2, 10, 5, 1}
	min, max := getMinMax(arr)
	fmt.Println("最小值：", min)
	fmt.Println("最大值：", max)
}

func getMinMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}