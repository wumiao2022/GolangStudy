package main

import "fmt"

func main() {
	// 练习1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3: 判断一个数是否为奇数
	num := 7
	if num%2 == 1 {
		fmt.Println("奇数")
	} else {
		fmt.Println("偶数")
	}

	// 练习4: 定义一个数组，填充1到5的数字，然后打印出来
	arr := [5]int{1, 2, 3, 4, 5}
	for _, v := range arr {
		fmt.Println(v)
	}

	// 练习5: 使用函数交换两个变量的值
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}

// 交换两个变量的值
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
