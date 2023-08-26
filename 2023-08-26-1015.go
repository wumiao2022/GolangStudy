package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个整数的和
	a := 10
	b := 5
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：计算两个浮点数的乘积
	x := 3.14
	y := 2.5
	product := x * y
	fmt.Println("Product:", product)

	// 练习4：判断一个数是否是偶数
	num := 20
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习5：循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6：使用数组保存一组字符串，并打印出所有元素
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// 练习7：使用map保存学生的成绩，并打印出所有学生的成绩
	scores := map[string]int{
		"Tom":   80,
		"Jerry": 90,
		"Alice": 95,
	}
	for name, score := range scores {
		fmt.Println(name, "score:", score)
	}
}