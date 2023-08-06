package main

import "fmt"

func main() {
	// 1. 定义一个整型数组并初始化
	numbers := [5]int{1, 2, 3, 4, 5}

	// 2. 遍历数组并打印每个元素
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// 3. 定义一个字符串切片并初始化
	names := []string{"Alice", "Bob", "Charlie", "Dave"}

	// 4. 遍历切片并打印每个元素
	for _, name := range names {
		fmt.Println(name)
	}

	// 5. 定义一个二维整型切片并初始化
	matrix := [][]int{{1, 2}, {3, 4}, {5, 6}}

	// 6. 遍历二维切片并打印每个元素
	for _, row := range matrix {
		for _, num := range row {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}

	// 7. 定义一个map并初始化
	ages := map[string]int{
		"Alice":  25,
		"Bob":    30,
		"Charlie": 35,
	}

	// 8. 遍历map并打印每个键值对
	for name, age := range ages {
		fmt.Println(name, age)
	}
}