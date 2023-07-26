package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义一个整数变量x，并将其初始化为10，然后打印出来
	x := 10
	fmt.Println(x)

	// 3. 定义一个字符串变量name，并将其初始化为"Go"，然后打印出来
	name := "Go"
	fmt.Println(name)

	// 4. 定义一个函数add，接收两个整数参数a和b，并返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 5))

	// 5. 定义一个切片slice，包含整数1、2、3，然后遍历输出每个元素
	slice := []int{1, 2, 3}
	for _, num := range slice {
		fmt.Println(num)
	}

	// 6. 使用for循环打印出0到9的数字
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 7. 定义一个结构体person，包含name和age字段，并初始化一个person对象
	type person struct {
		name string
		age  int
	}
	p := person{
		name: "Alice",
		age:  25,
	}
	fmt.Println(p)

	// 8. 使用if语句判断一个数是否为偶数，如果是偶数则打印"Even"，否则打印"Odd"
	num := 6
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 9. 使用switch语句判断一个数的值，如果是0则打印"Zero"，如果是1则打印"One"，否则打印"Other"
	num = 1
	switch num {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	default:
		fmt.Println("Other")
	}

	// 10. 使用defer关键字延迟执行一个打印语句
	defer fmt.Println("Deferred Print")

	// 11. 使用panic函数引发一个panic错误，并使用recover函数恢复错误
	panic("Panic Error")
	recover()
}