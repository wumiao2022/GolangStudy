package main

import "fmt"

func main() {

	// 1. 打印 "Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 变量声明和初始化
	var num1 int = 10
	var num2 int = 5
	var result int

	// 3. 变量计算
	result = num1 + num2

	// 4. 打印结果
	fmt.Println("结果为：", result)

	// 5. 条件语句
	if result > 10 {
		fmt.Println("结果大于10")
	} else {
		fmt.Println("结果小于等于10")
	}

	// 6. 循环语句
	for i := 0; i < 5; i++ {
		fmt.Println("当前循环次数：", i)
	}

	// 7. 数组声明和遍历
	numbers := [5]int{1, 2, 3, 4, 5}
	for i, value := range numbers {
		fmt.Println("当前索引：", i, "，对应的数值为：", value)
	}

	// 8. 切片声明和遍历
	slice := []int{1, 2, 3, 4, 5}
	for i, value := range slice {
		fmt.Println("当前索引：", i, "，对应的数值为：", value)
	}

	// 9. 函数声明和调用
	result := sum(10, 20)
	fmt.Println("函数调用结果为：", result)
}

// 10. 函数定义
func sum(a, b int) int {
	return a + b
}