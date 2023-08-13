package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：声明一个整数变量并打印其值
	var num1 int
	fmt.Println(num1)

	// 练习3：声明一个字符串变量并打印其值
	var str1 string
	fmt.Println(str1)

	// 练习4：声明一个浮点数变量并打印其值
	var num2 float64
	fmt.Println(num2)

	// 练习5：声明一个布尔类型变量并打印其值
	var isTrue bool
	fmt.Println(isTrue)

	// 练习6：声明一个切片并打印其长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// 练习7：声明一个函数并调用它
	myFunc := func() {
		fmt.Println("Calling myFunc")
	}
	myFunc()

	// 练习8：声明一个结构体并打印其字段的值
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"Alice", 20}
	fmt.Println(p.Name, p.Age)

	// 练习9：使用for循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习10：使用if语句判断一个数是否为偶数
	num3 := 6
	if num3%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}