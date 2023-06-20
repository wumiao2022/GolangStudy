package main

import "fmt"

func main() {
	// 练习题1：打印所有小于100的质数
	for i := 2; i < 100; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}

	// 练习题2：实现冒泡排序
	arr := []int{5, 3, 8, 2, 7, 1}
	bubbleSort(arr)
	fmt.Println(arr)

	// 练习题3：实现函数查找指定字符串在切片中的位置,若不存在则返回-1
	slice := []string{"apple", "banana", "orange", "grape"}
	fmt.Println(findStringIndex(slice, "orange"))
	fmt.Println(findStringIndex(slice, "watermelon"))
}

// 判断是否为质数
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// 冒泡排序
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 查找字符串在切片中的位置
func findStringIndex(slice []string, str string) int {
	for i := range slice {
		if slice[i] == str {
			return i
		}
	}
	return -1
}