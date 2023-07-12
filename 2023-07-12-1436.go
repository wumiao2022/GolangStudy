package main

import "fmt"

func main() {
	// 练习1：打印出Hello World!
	fmt.Println("Hello World!")

	// 练习2：创建一个整数变量并赋值为10，打印出该变量的值
	num := 10
	fmt.Println(num)

	// 练习3：创建一个字符串变量并赋值为"Hello"，再创建一个字符串变量并赋值为"World"，将这两个字符串拼接起来并打印出来
	str1 := "Hello"
	str2 := "World"
	fmt.Println(str1 + str2)

	// 练习4：创建一个整数数组并初始化为1、2、3、4、5，打印出数组中的第三个元素
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums[2])

	// 练习5：创建一个map，其中键为字符串类型，值为整数类型，将键值对分别赋值为"apple"和10，然后打印map的内容
	myMap := make(map[string]int)
	myMap["apple"] = 10
	fmt.Println(myMap)

	// 练习6：使用for循环打印出1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习7：创建一个函数，接收两个整数参数并返回它们的和
	sum := func(a, b int) int {
		return a + b
	}

	// 测试函数
	fmt.Println(sum(5, 3))
}