package main

import (
	"fmt"
)

func main() {
	// 练习1: 定义一个整型变量，将其初始化为10，并打印该变量的值
	var num1 int = 10
	fmt.Println(num1)

	// 练习2: 定义一个整型变量，并使用短变量声明将其初始化为20，并打印该变量的值
	num2 := 20
	fmt.Println(num2)

	// 练习3: 定义一个字符串变量，将其初始化为"Hello, Golang!"，并打印该变量的值
	var str1 string = "Hello, Golang!"
	fmt.Println(str1)

	// 练习4: 使用字面量定义一个布尔型变量，将其初始化为true，并打印该变量的值
	bool1 := true
	fmt.Println(bool1)

	// 练习5: 定义一个切片，元素类型为整型，将其初始化为[1, 2, 3, 4, 5]，并打印该切片的值
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
}