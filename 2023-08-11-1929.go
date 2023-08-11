package main

import (
	"fmt"
	"strings"
)

func main() {
	// 练习1: 将字符串中的某个字符替换为另一个字符
	str := "Hello, World!"
	str = strings.ReplaceAll(str, "o", "0")
	fmt.Println(str)

	// 练习2: 判断一个字符串是否为回文字符串
	str2 := "level"
	isPalindrome := true
	for i := 0; i < len(str2)/2; i++ {
		if str2[i] != str2[len(str2)-i-1] {
			isPalindrome = false
			break
		}
	}
	fmt.Println(isPalindrome)

	// 练习3: 在切片中查找某个元素的索引位置
	nums := []int{1, 2, 3, 4, 5}
	target := 3
	found := false
	for i, num := range nums {
		if num == target {
			fmt.Println("Found at index:", i)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Not found")
	}
}
