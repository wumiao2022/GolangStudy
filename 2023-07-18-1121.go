package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World!")

	// 练习2：计算两个数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数
	num := 8
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习4：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个结构体并初始化
	type person struct {
		name string
		age  int
	}
	p1 := person{name: "John", age: 30}
	fmt.Println(p1)

	// 练习6：使用defer关键字延迟执行函数
	defer fmt.Println("Deferred execution")
	fmt.Println("Executed first")
}