package main

import (
	"fmt"
)

func main() {
	// 练习1：输出1-10之间的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}

	// 练习2：计算1到100之间所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习3：将字符串反转并输出
	str := "hello world"
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()

	// 练习4：使用switch语句判断一个数字是否为奇数或偶数
	num := 5
	switch num % 2 {
	case 0:
		fmt.Println("偶数")
	case 1:
		fmt.Println("奇数")
	}

	// 练习5：定义一个结构体表示一个人，包含姓名、年龄、性别三个属性，并输出
	type person struct {
		name   string
		age    int
		gender string
	}
	p := person{"张三", 20, "男"}
	fmt.Println(p)
}