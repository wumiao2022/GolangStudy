package main

import "fmt"

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：变量赋值与输出
	var x int = 5
	var y float64 = 10.5
	fmt.Println("x =", x)
	fmt.Println("y =", y)

	// 练习3：条件判断与循环
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "是偶数")
		} else {
			fmt.Println(i, "是奇数")
		}
	}

	// 练习4：切片和循环遍历
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Println("索引", index, "的值是", value)
	}

	// 练习5：函数定义和调用
	result := add(2, 3)
	fmt.Println("2 + 3 =", result)
}

// 函数：求和
func add(a, b int) int {
	return a + b
}