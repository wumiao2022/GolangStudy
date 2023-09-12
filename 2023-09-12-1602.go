package main

import "fmt"

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：声明并初始化一个整型变量，并将其输出
	num := 10
	fmt.Println(num)

	// 练习3：使用if条件语句判断一个数是否为偶数，并输出结果
	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// 练习4：使用for循环输出1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个切片，将1到5的整数添加到切片中，并输出切片的长度和元素
	var slice []int
	for i := 1; i <= 5; i++ {
		slice = append(slice, i)
	}
	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice elements:", slice)
}