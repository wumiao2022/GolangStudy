package main

import "fmt"

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 定义一个变量并打印该变量的值
	var num int = 10
	fmt.Println(num)

	// 练习3: 使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4: 判断一个数是否为奇数
	num := 5
	if num%2 == 1 {
		fmt.Println("是奇数")
	} else {
		fmt.Println("不是奇数")
	}

	// 练习5: 定义一个切片并遍历该切片的元素
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}
}
