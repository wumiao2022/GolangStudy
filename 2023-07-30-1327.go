package main

import "fmt"

func main() {
	// 1. 求两个整数之和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 2. 判断一个数是奇数还是偶数
	num := 21
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 3. 将字符串转换为整数
	str := "123"
	num, _ = strconv.Atoi(str)
	fmt.Println("Number:", num)

	// 4. 交换两个变量的值
	a := 5
	b := 10

	// 方法1: 使用中间变量
	temp := a
	a = b
	b = temp
	fmt.Println("Swapped Values (Method 1):", a, b)

	// 方法2: 不使用中间变量
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("Swapped Values (Method 2):", a, b)

	// 5. 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, i*j)
		}
		fmt.Println()
	}
}