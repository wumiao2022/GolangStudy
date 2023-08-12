package main

import "fmt"

func main() {
	// 练习1: 定义一个变量x，并将其初始化为42，打印出x的值
	var x int = 42
	fmt.Println(x)

	// 练习2: 定义一个常量pi，并将其赋值为3.14，打印出pi的值
	const pi float64 = 3.14
	fmt.Println(pi)

	// 练习3: 通过循环打印出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4: 定义一个数组myArray，包含元素1、2、3、4、5，然后通过循环打印出数组元素
	myArray := [5]int{1, 2, 3, 4, 5}
	for _, val := range myArray {
		fmt.Println(val)
	}

	// 练习5: 定义一个切片mySlice，包含元素6、7、8、9、10，然后通过循环打印出切片元素
	mySlice := []int{6, 7, 8, 9, 10}
	for _, val := range mySlice {
		fmt.Println(val)
	}
}