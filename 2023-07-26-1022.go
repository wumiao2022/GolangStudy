package main

import "fmt"

func main() {
	// 练习1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 定义一个变量，赋值为10，然后输出该变量的值
	var num int = 10
	fmt.Println(num)

	// 练习3: 定义一个字符串变量，并给它赋值为"Go语言编程"
	var str string = "Go语言编程"
	fmt.Println(str)

	// 练习4: 定义一个整型数组，包含5个元素，然后打印数组的长度和第三个元素的值
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr))
	fmt.Println(arr[2])

	// 练习5: 使用for循环输出1到10之间的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6: 使用if语句判断一个数是否为正数，并输出判断结果
	num = -5
	if num > 0 {
		fmt.Println("该数是正数")
	} else {
		fmt.Println("该数不是正数")
	}
}