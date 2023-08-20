package main

import (
	"fmt"
)

func main() {
	// 1. 定义一个变量x，并赋值为10
	x := 10

	// 2. 定义一个常量y，并赋值为5
	const y = 5

	// 3. 打印变量x和常量y的值
	fmt.Println("x =", x)
	fmt.Println("y =", y)

	// 4. 定义一个切片a，并初始化为包含元素1、2、3的切片
	a := []int{1, 2, 3}

	// 5. 使用for循环遍历切片a，并打印每个元素的值
	for i := 0; i < len(a); i++ {
		fmt.Println("a[", i, "] =", a[i])
	}

	// 6. 定义一个映射b，并初始化为包含键值对"a": 1, "b": 2的映射
	b := map[string]int{
		"a": 1,
		"b": 2,
	}

	// 7. 使用for range循环遍历映射b，并打印每个键值对的值
	for key, value := range b {
		fmt.Println(key, ":", value)
	}
}
