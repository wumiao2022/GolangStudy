package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 练习1：生成一个长度为10的随机整数数组，并输出数组的元素
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	// 练习2：计算1到100之间所有奇数的和，并输出结果
	sum := 0
	for i := 1; i <= 100; i += 2 {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：编写一个函数，接收一个字符串，返回该字符串的长度和首字母大写后的字符串
	str := "hello world"
	length := len(str)
	capitalizedStr := string(str[0]-'a'+'A') + str[1:]
	fmt.Println(length, capitalizedStr)
}