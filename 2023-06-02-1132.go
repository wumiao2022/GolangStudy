package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 练习1：生成一个随机整数数组，并打印出来
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	// 练习2：将数组按照从大到小排序，并打印出来
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

	// 练习3：计算数组中所有元素的平均值
	sum := 0
	for _, v := range arr {
		sum += v
	}
	avg := float64(sum) / float64(len(arr))
	fmt.Println(avg)

	// 练习4：将平均值四舍五入并打印出来
	roundAvg := int(avg + 0.5)
	fmt.Println(roundAvg)

	// 练习5：将所有元素的值乘以平均值，并打印出来
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * roundAvg
	}
	fmt.Println(arr)
}