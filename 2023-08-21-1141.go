package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1: 输出当前时间
	fmt.Println("当前时间:", time.Now())

	// 练习2: 声明一个整数变量x并初始化为10，打印输出x的值
	x := 10
	fmt.Println("x的值为:", x)

	// 练习3: 声明一个字符串变量name并初始化为"John"，打印输出name的值
	name := "John"
	fmt.Println("name的值为:", name)

	// 练习4: 声明一个浮点数变量pi并初始化为3.14，打印输出pi的值
	pi := 3.14
	fmt.Println("pi的值为:", pi)

	// 练习5: 声明一个布尔变量isTrue并初始化为true，打印输出isTrue的值
	isTrue := true
	fmt.Println("isTrue的值为:", isTrue)

	// 练习6: 声明一个数组arr，包含元素1, 2, 3, 4, 5并打印输出数组中每个元素的值
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("数组arr的元素分别为:")
	for _, num := range arr {
		fmt.Println(num)
	}

	// 练习7: 声明一个切片slice，包含元素"A", "B", "C"并打印输出切片中每个元素的值
	slice := []string{"A", "B", "C"}
	fmt.Println("切片slice的元素分别为:")
	for _, letter := range slice {
		fmt.Println(letter)
	}

	// 练习8: 声明一个映射map，包含键值对"one":1, "two":2, "three":3并打印输出每个键和对应的值
	mp := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("映射mp的键值对为:")
	for key, value := range mp {
		fmt.Println(key, ":", value)
	}

	// 练习9: 声明一个指针变量ptr，指向整数变量x，并打印输出ptr的值和指针指向的变量的值
	ptr := &x
	fmt.Println("指针ptr的值为:", ptr)
	fmt.Println("指针ptr指向的变量的值为:", *ptr)

	// 练习10: 声明一个匿名函数，将其赋值给变量fn，调用fn并打印结果
	fn := func() {
		fmt.Println("这是一个匿名函数")
	}
	fn()
}