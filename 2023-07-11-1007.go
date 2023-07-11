package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 输入两个整数，求其和并输出
	var num1, num2 int
	fmt.Println("请输入两个整数：")
	fmt.Scanln(&num1, &num2)
	sum := num1 + num2
	fmt.Println("两个整数的和为：", sum)

	// 3. 定义一个字符串变量，并输出其长度
	str := "Golang"
	fmt.Println("字符串的长度为：", len(str))

	// 4. 定义一个整型切片，添加元素后打印切片长度和容量
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("切片的长度为：", len(slice))
	fmt.Println("切片的容量为：", cap(slice))

	// 5. 定义一个字典，添加键值对后打印字典长度
	dict := make(map[string]int)
	dict["one"] = 1
	dict["two"] = 2
	fmt.Println("字典的长度为：", len(dict))
}