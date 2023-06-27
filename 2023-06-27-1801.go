package main

import "fmt"

func main() {
	// 1. 输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 变量赋值及输出
	var a int = 10
	var b int = 20
	var c int

	c = a + b

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	// 3. 条件语句及输出
	if c > 30 {
		fmt.Printf("c 大于 30\n")
	} else {
		fmt.Printf("c 小于等于 30\n")
	}

	// 4. 循环语句及输出
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	// 5. 函数定义及调用
	fmt.Printf("5 加 7 的和为： %d\n", add(5, 7))
}

// 定义一个加法函数
func add(x int, y int) int {
	return x + y
}