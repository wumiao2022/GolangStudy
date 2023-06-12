package main

import (
	"fmt"
)

func main() {
	// 练习一：判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("这是一个偶数")
	} else {
		fmt.Println("这不是一个偶数")
	}

	// 练习二：将一个整数转换成二进制表示
	num2 := 18
	fmt.Printf("%b\n", num2)

	// 练习三：判断一个字符串是否为空
	str := " "
	if str == "" {
		fmt.Println("这是一个空字符串")
	} else {
		fmt.Println("这不是一个空字符串")
	}

	// 练习四：获取一个数组中的最大值
	arr := []int{1, 3, 5, 7, 9}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println("最大值为：", max)
}