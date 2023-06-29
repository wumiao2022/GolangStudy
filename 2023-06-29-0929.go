package main

import (
	"fmt"
)

func main() {
	// 练习1：声明一个字符串变量，值为 "Hello, World!"，并打印出来

	var str string = "Hello, World!"
	fmt.Println(str)

	// 练习2：声明一个整数变量，值为 100，打印出来

	var num int = 100
	fmt.Println(num)

	// 练习3：声明一个浮点数变量，值为 3.14，打印出来

	var floatNum float64 = 3.14
	fmt.Println(floatNum)

	// 练习4：声明一个布尔型变量，值为 true，打印出来

	var boolVar bool = true
	fmt.Println(boolVar)

	// 练习5：声明一个切片，包含整数元素 1、2、3、4、5，并打印出来

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 练习6：声明一个字典，包含键值对 "name": "John"，并打印出来

	dict := map[string]string{"name": "John"}
	fmt.Println(dict)

	// 练习7：声明一个函数，接受两个整数参数，并返回它们的和

	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum(2, 3))
}