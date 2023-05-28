package main

import "fmt"

func main() {
	// 练习1：输出 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算 1+2+...+10 的结果并输出
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：求出数组 [1, 2, 3, 4, 5] 的平均值并输出
	arr := [5]int{1, 2, 3, 4, 5}
	sum = 0
	for _, value := range arr {
		sum += value
	}
	avg := float64(sum) / float64(len(arr))
	fmt.Println(avg)

	// 练习4：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}

	// 练习5：使用函数交换两个变量的值
	a, b := 1, 2
	swap(&a, &b)
	fmt.Printf("a=%d, b=%d\n", a, b)
}

func swap(x, y *int) {
	*x, *y = *y, *x
}