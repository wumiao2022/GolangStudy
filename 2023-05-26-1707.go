package main

import (
	"fmt"
)

func main() {
	// 示例1：计算1-100之间所有奇数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			sum += i
		}
	}
	fmt.Println("1-100之间所有奇数的和为:", sum)

	// 示例2：找出数组中的最大值和最小值
	nums := []int{12, 45, 64, 22, 7, 98, 31, 56}
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
	fmt.Printf("数组%v中的最大值为%d，最小值为%d\n", nums, max, min)

	// 示例3：将一个整数反转
	num := 123456789
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	fmt.Printf("反转后的整数为:%d\n", reversed)
}