package main

import "fmt"

func main() {
	// 实例1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 实例2：计算两个数的和
	a := 10
	b := 5
	sum := a + b
	fmt.Println("Sum:", sum)

	// 实例3：打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 实例4：判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// 实例5：使用函数交换两个变量的值
	x := 10
	y := 5
	swap(&x, &y)
	fmt.Println("After swap, x =", x, "y =", y)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
