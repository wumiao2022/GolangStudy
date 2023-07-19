package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量和类型
	var num1 int = 10
	var num2 float64 = 3.14
	var str string = "Golang"
	fmt.Println(num1, num2, str)

	// 练习3：基本运算
	var a int = 10
	var b int = 5
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// 练习4：条件和循环
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// 练习5：数组和切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)

	// 练习6：结构体和方法
	type Person struct {
		Name   string
		Age    int
		Gender string
	}
	p := Person{Name: "Alice", Age: 25, Gender: "Female"}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Gender:", p.Gender)

	// 练习7：并发编程
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	fmt.Println(<-ch)
}