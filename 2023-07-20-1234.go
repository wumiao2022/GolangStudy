package main

import "fmt"

func main() {
	// 1. 声明一个整型变量num并初始化为10，打印出num的值
	num := 10
	fmt.Println(num)

	// 2. 声明一个字符串变量name并初始化为空字符串，打印出name的值
	var name string
	fmt.Println(name)

	// 3. 声明一个切片变量colors并初始化为包含红、绿、蓝三个元素的字符串切片，打印出colors的长度和第二个元素的值
	colors := []string{"红", "绿", "蓝"}
	fmt.Println(len(colors), colors[1])

	// 4. 声明一个映射变量person并初始化为包含姓名、年龄的键值对，姓名为"John"，年龄为25，打印出person的值
	person := map[string]int{"姓名": 25}
	fmt.Println(person)

	// 5. 声明一个函数add，接受两个整型参数，并返回它们的和。调用add函数并打印出结果
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(5, 3))
}