package main

import "fmt"

func main() {
	// 练习：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习：计算并打印两个整数的和
	var num1, num2 int = 5, 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习: 使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习：使用if-else判断一个整数是奇数还是偶数，并打印结果
	num := 8
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习：创建一个切片，并打印切片的长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// 练习：使用switch-case判断一个数字是正数、负数还是零，并打印结果
	number := -5
	switch {
	case number > 0:
		fmt.Println("Positive")
	case number < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}

	// 练习：定义一个结构体，并初始化其字段，并打印该结构体的值
	type person struct {
		name string
		age  int
	}
	p := person{name: "Alice", age: 25}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

运行结果：
Hello, World!
Sum: 15
1
2
3
4
5
6
7
8
9
10
Even
Length: 5
Capacity: 5
Negative
Name: Alice
Age: 25