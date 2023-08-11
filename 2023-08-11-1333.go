package main

import (
	"fmt"
)

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义一个变量x为整数类型，赋值为10，并输出x的值
	x := 10
	fmt.Println(x)

	// 3. 定义一个字符串变量name，赋值为你的名字，并以“Hello, [name]!”形式输出
	name := "Alice"
	fmt.Printf("Hello, %s!\n", name)

	// 4. 定义一个切片slice，包含整数元素1, 2, 3, 4, 5，并输出切片的长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// 5. 使用for循环打印出切片slice中的每一个元素
	for _, value := range slice {
		fmt.Println(value)
	}

	// 6. 定义一个函数add，接受两个整数参数并返回它们的和，在main函数中调用add函数并输出结果
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))

	// 7. 使用if语句判断一个数num是否为偶数，如果是则输出"num is even"，否则输出"num is odd"
	num := 7
	if num%2 == 0 {
		fmt.Println("num is even")
	} else {
		fmt.Println("num is odd")
	}
}