package main

import (
	"fmt"
)

func main() {
	// 练习1：将整型数组{1, 2, 3, 4, 5}的所有元素乘以2并打印出来
	nums := []int{1, 2, 3, 4, 5}
	for i := range nums {
		nums[i] *= 2
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()

	// 练习2：从控制台输入一个字符串，将其中的每个字符打印出来
	var str string
	fmt.Scanln(&str)
	for _, ch := range str {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 练习3：将数组{"apple", "banana", "orange"}中元素的首字母大写并打印出来
	strs := []string{"apple", "banana", "orange"}
	for _, str := range strs {
		fmt.Printf("%c%s ", 'A'+str[0]-'a', str[1:])
	}
	fmt.Println()

	// 练习4：将整数123456789的每位数字打印出来
	num := 123456789
	for num > 0 {
		fmt.Printf("%d ", num % 10)
		num /= 10
	}
	fmt.Println()
}