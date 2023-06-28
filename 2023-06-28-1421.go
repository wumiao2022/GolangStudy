package main

import "fmt"

func main() {
	// 练习1：输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：创建一个切片，包含数字1到5，然后打印出切片中的所有元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习3：创建一个函数，接受一个整数参数x，并返回其平方值
	fmt.Println(square(5))

	// 练习4：使用for循环打印出1到10之间的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习5：使用if语句判断一个年份是否是闰年，如果是则输出"是闰年"，否则输出"不是闰年"
	year := 2024
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println("是闰年")
	} else {
		fmt.Println("不是闰年")
	}
}

func square(x int) int {
	return x * x
}
