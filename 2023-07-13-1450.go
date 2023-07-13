package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印"Hello, world!"
	fmt.Println("Hello, world!")

	// 练习2: 声明一个整数变量并赋值为10，输出变量的值
	num := 10
	fmt.Println(num)

	// 练习3: 声明一个字符串变量并赋值为"Golang"，输出变量的值
	str := "Golang"
	fmt.Println(str)

	// 练习4: 声明一个布尔变量并赋值为true，输出变量的值
	flag := true
	fmt.Println(flag)
	
	// 练习5：声明一个数组变量arr，包含[1, 2, 3, 4, 5]五个整数，并分别输出数组的每一个元素
	arr := [5]int{1, 2, 3, 4, 5}
	for _, value := range arr {
		fmt.Println(value)
	}
	
	// 练习6：声明一个切片变量slice，包含[6, 7, 8, 9, 10]五个整数，并分别输出切片的每一个元素
	slice := []int{6, 7, 8, 9, 10}
	for _, value := range slice {
		fmt.Println(value)
	}
	
	// 练习7：使用for循环输出1到10的所有整数（包含1和10）
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	
	// 练习8：使用if语句判断一个整数num是奇数还是偶数，并输出判断结果
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	
	// 练习9：使用switch语句根据一个星期的天数（1-7）输出对应的星期几
	day := 5
	switch day {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("输入错误")
	}
}