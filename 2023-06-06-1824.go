package main

import (
	"fmt"
)

func main() {
	// 练习1：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习2：统计一个字符串中每个字母出现的次数
	str := "hello world"
	countMap := make(map[rune]int)

	for _, c := range str {
		countMap[c]++
	}

	for k, v := range countMap {
		fmt.Printf("%c：%d\n", k, v)
	}

	// 练习3：数组反转
	arr := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr)/2; i++ {
		tmp := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = tmp
	}
	fmt.Println(arr)
}