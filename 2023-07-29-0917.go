package main

import "fmt"

func main() {
	// 1. 交换两个变量的值
	a := 10
	b := 20
	a, b = b, a
	fmt.Println("a:", a, "b:", b)

	// 2. 判断一个数是奇数还是偶数
	num := 11
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 3. 遍历切片并打印元素
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// 4. 使用递归计算斐波那契数列的第n个数
	n := 5
	fmt.Println(fibonacci(n))

	// 5. 使用递归反转字符串
	str := "hello"
	fmt.Println(reverseString(str))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func reverseString(str string) string {
	if len(str) <= 1 {
		return str
	} else {
		return reverseString(str[1:]) + str[:1]
	}
}