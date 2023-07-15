package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：声明一个整数变量并给其赋值为10，然后打印出该变量的值
	var num int = 10
	fmt.Println(num)

	// 练习3：声明一个字符串变量并给其赋值为"Go语言是最好的编程语言"，然后打印出该变量的值
	var str string = "Go语言是最好的编程语言"
	fmt.Println(str)

	// 练习4：声明一个布尔变量并给其赋值为true，然后打印出该变量的值
	var flag bool = true
	fmt.Println(flag)

	// 练习5：声明一个浮点数变量并给其赋值为3.14，然后打印出该变量的值
	var pi float64 = 3.14
	fmt.Println(pi)

	// 练习6：声明一个整数数组，并使用循环打印数组中的元素
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 练习7：声明一个切片，并使用循环打印切片中的元素
	slice := []int{1, 2, 3, 4, 5}
	for _, val := range slice {
		fmt.Println(val)
	}

	// 练习8：声明一个映射，并使用循环打印映射中的键值对
	m := map[string]int{"apple": 1, "banana": 2, "orange": 3}
	for key, val := range m {
		fmt.Println(key, val)
	}
}