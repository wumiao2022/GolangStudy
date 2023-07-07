package main

import "fmt"

func main() {
	// 1. 打印"Hello, World!"到控制台
	fmt.Println("Hello, World!")

	// 2. 定义一个整型变量x，并赋值为10，打印x的值
	x := 10
	fmt.Println(x)

	// 3. 定义一个字符串变量name，并赋值为"GoLang"，打印name的值
	name := "GoLang"
	fmt.Println(name)

	// 4. 定义一个bool类型的变量isTrue并赋值为true，打印isTrue的值
	isTrue := true
	fmt.Println(isTrue)

	// 5. 定义一个切片变量numbers，并初始化为包含1, 2, 3三个元素的切片，打印numbers的值
	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	// 6. 定义一个常量pi，并赋值为3.14159，打印pi的值
	const pi = 3.14159
	fmt.Println(pi)

	// 7. 定义一个函数add，接收两个整型参数x和y，返回它们的和
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(2, 3))

	// 8. 使用for循环打印从1到10的所有整数
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 9. 使用if语句判断一个整数是否为偶数，打印结果
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 10. 使用switch语句判断一个整数的值，并打印不同的结果
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}
}