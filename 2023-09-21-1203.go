package main

import (
	"fmt"
)

func main() {
	// 实例1：数组求和
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println("数组求和:", sum)

	// 实例2：切片反转
	slice := []string{"apple", "banana", "orange", "grape"}
	reverse := make([]string, len(slice))
	for i, j := 0, len(slice)-1; i < len(slice); i, j = i+1, j-1 {
		reverse[j] = slice[i]
	}
	fmt.Println("切片反转:", reverse)

	// 实例3：map遍历
	mp := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 3,
		"grape":  7,
	}
	fmt.Println("map遍历:")
	for key, value := range mp {
		fmt.Printf("%s: %d\n", key, value)
	}
}