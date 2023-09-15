package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量和常量
	var a int = 10
	fmt.Println(a)

	const b string = "Golang"
	fmt.Println(b)

	// 练习3：数组和切片
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 练习4：循环和条件判断
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习5：函数和闭包
	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println(sum(5, 3))

	// 练习6：指针和结构体
	type Person struct {
		Name string
		Age  int
	}

	p := &Person{
		Name: "Alice",
		Age:  20,
	}

	fmt.Println(p)

	// 练习7：并发和通道
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	fmt.Println(<-ch)
}