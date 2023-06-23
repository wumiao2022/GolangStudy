package main

import (
	"fmt"
)

func main() {
	// 练习1：计算两个整数之和
	var a, b int = 10, 20
	fmt.Println("a + b =", a+b)

	// 练习2：判断一个数是偶数还是奇数
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 练习3：将一个字符串反转
	str := "hello world"
	result := []rune(str)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	fmt.Println(string(result))

	// 练习4：遍历数组，将数组元素翻倍并打印出来
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i] * 2)
	}
	
	// 练习5：遍历map，将key和value打印出来
	m := map[string]int{"apple": 1, "banana": 2, "orange": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}
}