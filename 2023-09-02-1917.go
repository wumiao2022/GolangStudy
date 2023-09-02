package main

import "fmt"

func main() {
	// 练习1: 打印Hello World
	fmt.Println("Hello World")

	// 练习2: 变量与常量
	var x int = 5
	const y float64 = 3.14
	fmt.Println("x =", x)
	fmt.Println("y =", y)

	// 练习3: 数组与切片
	arr := [3]int{1, 2, 3}
	slice := []int{4, 5, 6}
	fmt.Println("arr =", arr)
	fmt.Println("slice =", slice)

	// 练习4: 循环与条件语句
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// 练习5: 函数与方法
	result := add(3, 4)
	fmt.Println("3 + 4 =", result)

	// 练习6: 结构体与指针
	person := Person{Name: "Alice", Age: 25}
	fmt.Println(person.Name)
	fmt.Println(person.Age)

	// 练习7: 接口与类型断言
	var mobile Phone
	mobile = new(IPhone)
	mobile.call()
	switch mobile.(type) {
	case *IPhone:
		fmt.Println("This is an iPhone")
	default:
		fmt.Println("Unknown phone type")
	}
}

func add(a, b int) int {
	return a + b
}

type Person struct {
	Name string
	Age  int
}

type Phone interface {
	call()
}

type IPhone struct {
}

func (phone *IPhone) call() {
	fmt.Println("Calling from iPhone")
}
