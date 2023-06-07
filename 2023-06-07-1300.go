package main

import "fmt"

func main() {
	// 练习1：输出"Hello World"
	fmt.Println("Hello World")

	// 练习2：声明一个整型变量并赋值为10，输出这个变量的值
	var num int = 10
	fmt.Println(num)

	// 练习3：使用for循环输出1~10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：定义一个结构体类型Person，包含name和age两个字段
	type Person struct {
		name string
		age  int
	}

	// 练习5：定义一个包含5个元素的整型数组，输出数组中的每个元素值
	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 练习6：定义一个切片类型slice，并向其中添加5个元素，输出slice的长度和元素值
	var slice []int
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(len(slice), slice)

	// 练习7：定义一个map类型，包含3个键值对，输出map中的所有键值对
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}
}