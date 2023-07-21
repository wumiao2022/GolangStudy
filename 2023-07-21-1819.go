package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：声明一个整数变量并打印其值
	var num int
	fmt.Println(num)

	// 练习3：声明一个字符串变量并打印其值
	var str string
	fmt.Println(str)

	// 练习4：使用循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：声明一个切片并添加3个字符串元素
	slice := []string{"apple", "banana", "cherry"}
	fmt.Println(slice)

	// 练习6：声明一个map并添加3个键值对
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	m["cherry"] = 3
	fmt.Println(m)
}