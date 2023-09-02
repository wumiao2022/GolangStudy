package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量num，并赋值为10
	num := 10

	// 3. 声明一个字符串变量name，并赋值为"GoLang"
	name := "GoLang"

	// 4. 打印变量num和name的值
	fmt.Println("num:", num)
	fmt.Println("name:", name)

	// 5. 定义一个切片slice1，包含元素1, 2, 3, 4, 5
	slice1 := []int{1, 2, 3, 4, 5}

	// 6. 打印切片slice1的长度和容量
	fmt.Println("slice1 length:", len(slice1))
	fmt.Println("slice1 capacity:", cap(slice1))

	// 7. 声明一个map，键类型为string，值类型为int
	m := make(map[string]int)

	// 8. 向map中插入键值对
	m["apple"] = 5
	m["banana"] = 8

	// 9. 打印map的内容
	fmt.Println("map:", m)

	// 10. 使用循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}