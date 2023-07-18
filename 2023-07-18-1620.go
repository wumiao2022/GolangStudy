package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 实现一个函数，生成一个包含n个随机整数的切片，并返回该切片
	fmt.Println("Exercise 1:")
	fmt.Println(generateRandomSlice(10))

	// 实现一个函数，将一个切片中的所有元素进行翻转，并返回新的切片
	fmt.Println("Exercise 2:")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(reverseSlice(slice))

	// 实现一个函数，找出切片中的最大值和最小值，并返回它们
	fmt.Println("Exercise 3:")
	slice2 := []int{10, 5, 8, 2, 3, 9}
	max, min := findMinMax(slice2)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}

func generateRandomSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func reverseSlice(slice []int) []int {
	reversedSlice := make([]int, len(slice))
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		reversedSlice[i], reversedSlice[j] = slice[j], slice[i]
	}
	return reversedSlice
}

func findMinMax(slice []int) (int, int) {
	max, min := slice[0], slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
		if slice[i] < min {
			min = slice[i]
		}
	}
	return max, min
}