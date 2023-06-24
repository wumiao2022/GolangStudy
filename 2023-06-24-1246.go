package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 练习1：生成一个长度为10的随机整数数组，打印数组中的最大值、最小值和平均值
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("数组：", arr)
	min, max := arr[0], arr[0]
	sum := 0
	for i := 0; i < 10; i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
		sum += arr[i]
	}
	avg := float64(sum) / float64(10)
	fmt.Println("最小值：", min)
	fmt.Println("最大值：", max)
	fmt.Println("平均值：", avg)

	// 练习2：定义一个函数，输入一个字符串，反转字符串中的字符顺序后输出
	str := "hello world"
	fmt.Println(reverseStr(str))
}

func reverseStr(str string) string {
	s := []byte(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}