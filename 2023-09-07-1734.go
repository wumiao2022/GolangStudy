package main

import "fmt"

func main() {
	// 示例1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 示例2: 定义一个整型变量并赋值
	var num int = 10
	fmt.Println(num)

	// 示例3: 判断一个数是奇数还是偶数
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 示例4: 使用for循环计算1~10之间所有整数的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 示例5: 定义一个切片并遍历打印
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}

	// 示例6: 定义一个结构体类型并初始化
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 28}
	fmt.Println(p)
}