package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 定义一个变量并赋值为10，打印该变量的值
	num := 10
	fmt.Println(num)

	// 练习3: 定义一个函数add并传入两个整数参数，返回它们的和
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	// 练习4: 创建一个切片包含1到5的整数，然后打印切片中的所有元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习5: 创建一个映射，将城市名称映射到它们所在的国家
	cities := map[string]string{
		"Beijing":  "China",
		"Tokyo":    "Japan",
		"New York": "USA",
	}
	for city, country := range cities {
		fmt.Println(city, "is in", country)
	}
}

func add(a, b int) int {
	return a + b
}