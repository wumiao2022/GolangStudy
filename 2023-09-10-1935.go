package main

import "fmt"

func main() {
	// 1. 使用for循环输出1到10的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 2. 使用switch语句判断一个月份是属于哪个季节
	month := 3
	switch month {
	case 1, 2, 12:
		fmt.Println("冬季")
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	default:
		fmt.Println("无效的月份")
	}

	// 3. 使用递归函数计算斐波那契数列的第n项
	n := 7
	fmt.Printf("斐波那契数列的第%d项为：%d\n", n, fibonacci(n))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}