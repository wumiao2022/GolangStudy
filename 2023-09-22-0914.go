package main

import "fmt"

func main() {
	// 练习1：打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：定义一个整型变量x，并将其赋值为10
	x := 10

	// 练习3：定义一个字符串变量name，并将其赋值为"John"
	name := "John"

	// 练习4：定义一个浮点数变量y，并将其赋值为3.14
	y := 3.14

	// 练习5：创建一个切片a，包含整数1、2、3、4、5
	a := []int{1, 2, 3, 4, 5}

	// 练习6：创建一个map b，包含键值对 "apple": 3, "banana": 5, "orange": 2
	b := map[string]int{"apple": 3, "banana": 5, "orange": 2}

	// 练习7：编写一个for循环，将切片a中的元素依次打印出来
	for _, value := range a {
		fmt.Println(value)
	}

	// 练习8：编写一个if-else语句，判断x的大小，如果大于等于10则打印 "x is greater than or equal to 10"，否则打印 "x is less than 10"
	if x >= 10 {
		fmt.Println("x is greater than or equal to 10")
	} else {
		fmt.Println("x is less than 10")
	}

	// 练习9：编写一个函数add，接受两个整数参数，返回它们的和
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(4, 6))

	// 练习10：编写一个递归函数factorial，计算一个非负整数的阶乘
	factorial := func(n uint64) uint64 {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	}

	fmt.Println(factorial(5))
}