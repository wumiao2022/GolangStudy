package main

import (
	"fmt"
)

func main() {
	// 1. 交换两个变量的值
	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a) // 输出：20
	fmt.Println(b) // 输出：10

	// 2. 判断一个数是否为奇数
	num := 21
	if num%2 == 1 {
		fmt.Println("奇数")
	} else {
		fmt.Println("偶数")
	}

	// 3. 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}

	// 4. 判断一个字符串是否是回文字符串
	str := "level"
	isPalindrome := true
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("是回文字符串")
	} else {
		fmt.Println("不是回文字符串")
	}

	// 5. 实现冒泡排序
	arr := []int{5, 8, 1, 4, 2}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}