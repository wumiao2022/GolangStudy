package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义一个整数变量x，并赋值为42，并打印出来
	x := 42
	fmt.Println(x)

	// 3. 定义一个字符串变量name，并赋值为"John"，并打印出来
	name := "John"
	fmt.Println(name)

	// 4. 定义一个浮点数变量pi，并赋值为3.14，并打印出来
	pi := 3.14
	fmt.Println(pi)

	// 5. 定义一个布尔类型变量isTrue，并赋值为true，并打印出来
	isTrue := true
	fmt.Println(isTrue)

	// 6. 定义两个整数变量a和b，并分别赋值为10和5，然后求和并打印结果
	a := 10
	b := 5
	sum := a + b
	fmt.Println(sum)

	// 7. 使用for循环打印0到9的数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 8. 定义一个切片nums，包含整数1到5，并打印出来
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
}