package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// 生成一个随机数切片
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 10)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println("原始切片:", nums)

	// 获取切片中的最大值和最小值
	max := nums[0]
	min := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Println("最大值:", max)
	fmt.Println("最小值:", min)

	// 将切片中的偶数乘以2
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i] *= 2
		}
	}
	fmt.Println("修改后的切片:", nums)

	// 统计切片中奇数的个数
	count := 0
	for _, num := range nums {
		if num%2 != 0 {
			count++
		}
	}
	fmt.Println("奇数个数:", count)

	// 将切片中的元素转换为字符串并拼接起来
	str := ""
	for _, num := range nums {
		str += strconv.Itoa(num) + ","
	}
	fmt.Println("切片转换为字符串:", str)
}