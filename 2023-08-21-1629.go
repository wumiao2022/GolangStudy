package main

import "fmt"

func main() {
	// 练习1：打印出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和并输出结果
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并输出结果
	num := 17
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：定义一个切片并输出其中的元素
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 练习6：定义一个结构体并输出其中的字段值
	type Person struct {
		Name    string
		Age     int
		Country string
	}
	person := Person{Name: "John", Age: 25, Country: "USA"}
	fmt.Println(person)

	// 练习7：使用并发打印出1到10的数字
	for i := 1; i <= 10; i++ {
		go fmt.Println(i)
	}
	// 这里需要让主协程等待其他协程完成打印操作
	// 可以使用time.Sleep或者sync.WaitGroup实现
	// 这里简单使用time.Sleep演示
	time.Sleep(time.Second)
}