package main

import "fmt"

func main() {
	// 定义一个切片
	slice := []int{1, 2, 3, 4, 5}

	// 遍历切片，求和
	sum := 0
	for _, num := range slice {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// 将切片中的元素乘以2
	for i, num := range slice {
		slice[i] = num * 2
	}
	fmt.Println("Doubled Slice:", slice)

	// 定义一个映射
	mapping := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// 遍历映射
	for key, value := range mapping {
		fmt.Println(key, ":", value)
	}

	// 判断映射中是否存在某个键
	if value, ok := mapping["banana"]; ok {
		fmt.Println("Value of banana:", value)
	} else {
		fmt.Println("banana not found")
	}
}