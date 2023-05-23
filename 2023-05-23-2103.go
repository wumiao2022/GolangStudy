package main

import "fmt"

func main() {
	// 练习一：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习二：定义一个字符串变量并打印
	var str string = "Golang多练习"
	fmt.Println(str)

	// 练习三：定义一个整型变量并打印
	var num int = 100
	fmt.Println(num)

	// 练习四：定义一个浮点型变量并打印
	var flt float64 = 3.14
	fmt.Println(flt)

	// 练习五：定义一个布尔型变量并打印
	var b bool = true
	fmt.Println(b)

	// 练习六：定义一个切片并打印
	var arr []int = []int{1, 2, 3}
	fmt.Println(arr)

	// 练习七：定义一个映射并打印
	var m map[string]int = map[string]int{"A": 1, "B": 2}
	fmt.Println(m)
}