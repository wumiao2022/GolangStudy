package main

import (
	"fmt"
)

func main() {
	// 示例1-输出Hello World！
	fmt.Println("Hello World!")

	// 示例2-定义变量并赋值
	var a int = 10
	fmt.Println("a = ", a)

	// 示例3-流程控制for循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 示例4-切片
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// 示例5-函数
	sum := add(1, 2)
	fmt.Println("1 + 2 =", sum)
}

// add函数-求和
func add(x int, y int) int {
	return x + y
}