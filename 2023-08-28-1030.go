package main

import "fmt"

func main() {
	// 例子1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 例子2：计算两个整数的和
	a := 5
	b := 3
	sum := a + b
	fmt.Println("Sum:", sum)

	// 例子3：判断一个数是否为偶数
	num := 8
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 例子4：使用for循环打印1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 例子5：定义一个函数并调用
	greet("John")

	// 例子6：使用数组存储一组整数，并进行遍历
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 例子7：使用map存储字符串和对应的整数值
	ages := map[string]int{
		"John": 25,
		"Jane": 30,
	}
	fmt.Println(ages["John"])

	// 例子8：使用defer延迟执行一段代码
	defer fmt.Println("Deferred execution")

	// 例子9：使用接口实现多态
	shape1 := Circle{radius: 5}
	shape2 := Rectangle{width: 3, height: 4}

	shapes := []Shape{shape1, shape2}

	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
	}
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
