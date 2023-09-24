package main

import "fmt"

func main() {
	// 1. 打印 Hello, World!
	fmt.Println("Hello, World!")

	// 2. 声明并初始化一个整型变量 x，赋值为 5
	x := 5
	fmt.Println(x)

	// 3. 声明一个字符串变量 message，赋值为 "Hello, Golang!"
	var message string = "Hello, Golang!"
	fmt.Println(message)

	// 4. 定义一个数组 arr 包含 3 个整型元素，赋值为 1, 2, 3
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// 5. 定义一个切片 nums 包含 5 个整型元素，赋值为 1, 2, 3, 4, 5
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// 6. 利用 for 循环遍历切片 nums，并打印每个元素的值
	for _, num := range nums {
		fmt.Println(num)
	}

	// 7. 编写一个函数 add，接收两个整型参数，返回它们的和
	fmt.Println(add(3, 4))

	// 8. 编写一个结构体 Circle，有一个半径字段，并定义一个计算面积的方法
	c := Circle{radius: 5}
	fmt.Println(c.Area())

	// 9. 利用 defer 延迟打印 "Done"
	defer fmt.Println("Done")

	// 10. 使用 panic 和 recover 组合来处理异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered:", err)
		}
	}()
	panic("Oops, something went wrong!")
}

func add(a, b int) int {
	return a + b
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}