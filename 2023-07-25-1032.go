package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数
	randomNum := rand.Intn(100)
	fmt.Println("随机数:", randomNum)

	// 判断随机数的奇偶性
	if randomNum%2 == 0 {
		fmt.Println("随机数是偶数")
	} else {
		fmt.Println("随机数是奇数")
	}

	// 生成一个包含10个随机数的切片
	var randomSlice []int
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(100)
		randomSlice = append(randomSlice, randNum)
	}
	fmt.Println("随机数切片:", randomSlice)

	// 对切片进行排序
	for i := 0; i < len(randomSlice)-1; i++ {
		for j := 0; j < len(randomSlice)-i-1; j++ {
			if randomSlice[j] > randomSlice[j+1] {
				temp := randomSlice[j]
				randomSlice[j] = randomSlice[j+1]
				randomSlice[j+1] = temp
			}
		}
	}
	fmt.Println("排序后的随机数切片:", randomSlice)
}