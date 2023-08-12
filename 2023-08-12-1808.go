package main

import (
	"fmt"
	"time"
)

func main() {
	// 打印当前时间
	fmt.Println("Current Time:", time.Now())

	// 计算两个数的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 循环打印数字
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 判断两个数的关系
	x := 30
	y := 40
	if x > y {
		fmt.Println("x is greater than y")
	} else if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is equal to y")
	}

	// 定义并使用一个结构体
	type person struct {
		name string
		age  int
	}
	p := person{"Alice", 25}
	fmt.Println("Person:", p)

	// 使用切片
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// 使用map
	grades := map[string]int{"Math": 90, "English": 85, "Science": 95}
	fmt.Println("Grades:", grades)
}