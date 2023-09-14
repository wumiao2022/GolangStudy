package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 定义一个变量，初始化为10，打印出这个变量的值
	num := 10
	fmt.Println(num)

	// 练习3: 定义一个常量，值为3.14，打印出这个常量的值
	const PI = 3.14
	fmt.Println(PI)

	// 练习4: 定义一个数组，包含5个元素，分别是1, 2, 3, 4, 5，打印出数组的值
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 练习5: 定义一个切片，包含五个元素，分别是"a", "b", "c", "d", "e"，打印出切片的值
	slice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slice)
}