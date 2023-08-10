package main

import (
	"fmt"
)

// 1. 编写一个函数，接收一个整数切片作为参数，并返回该切片的长度和容量
func getSliceInfo(slice []int) (int, int) {
	return len(slice), cap(slice)
}

// 2. 编写一个函数，接收一个整数切片和一个整数作为参数，并将该整数添加到切片末尾，然后返回更新后的切片
func appendToSlice(slice []int, num int) []int {
	return append(slice, num)
}

// 3. 编写一个函数，接收一个整数slice作为参数，并返回该slice中所有元素的和
func sumOfSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

// 4. 编写一个函数，接收一个整数切片作为参数，并返回该切片中的最大值和最小值
func getMinMax(slice []int) (int, int) {
	min := slice[0]
	max := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// 5. 编写一个函数，接收两个整数切片作为参数，判断两个切片的内容是否完全相同
func compareSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func main() {
	// 测试函数
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("原始切片：", slice)

	length, capacity := getSliceInfo(slice)
	fmt.Println("切片长度：", length)
	fmt.Println("切片容量：", capacity)

	newSlice := appendToSlice(slice, 6)
	fmt.Println("更新后的切片：", newSlice)

	sum := sumOfSlice(slice)
	fmt.Println("切片元素和：", sum)

	min, max := getMinMax(slice)
	fmt.Println("切片最小值：", min)
	fmt.Println("切片最大值：", max)

	slice2 := []int{1, 2, 3, 4, 5}
	areEqual := compareSlices(slice, slice2)
	fmt.Println("两个切片是否相等：", areEqual)
}