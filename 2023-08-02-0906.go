package main

import "fmt"

func main() {
	// 练习1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和并打印结果
	a := 2
	b := 3
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)

	// 练习3：判断一个数是否为偶数，并打印结果
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}

	// 练习4：使用for循环打印1到10的平方数
	for i := 1; i <= 10; i++ {
		fmt.Println(i*i)
	}

	// 练习5：定义一个结构体类型，并创建一个结构体实例
	type Person struct {
		Name string
		Age  int
	}

	person := Person{
		Name: "Alice",
		Age:  25,
	}

	fmt.Println(person)

	// 练习6：定义一个切片，并使用append函数向切片中添加元素
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice)
}