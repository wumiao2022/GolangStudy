package main

import "fmt"

func main() {
	// 练习1：输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和并输出
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习3：使用for循环输出1到10的所有数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：定义一个数组并输出数组中的所有元素
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}

	// 练习5：定义一个函数，接收两个参数并返回它们的乘积
	multiply := func(a, b int) int {
		return a * b
	}
	product := multiply(3, 4)
	fmt.Println("The product is:", product)
}

// 练习6：定义一个结构体类型，并为其定义一个方法，用于打印该结构体的属性值
type Person struct {
	Name string
	Age  int
}

func (p Person) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
	p := Person{"John", 25}
	p.PrintInfo()
}