package main

import (
	"fmt"
)

func main() {
	// 练习1：打印所有小于1000的3或5的倍数的数字的和。
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

	// 练习2：编写一个函数，将一个字符串反转。
	str := "hello world"
	reverseStr := reverse(str)
	fmt.Println(reverseStr)

	// 练习3：将一维数组转换成二维数组（每行包含2个元素）。
	arr := []int{1, 2, 3, 4, 5, 6}
	twoDArray := toTwoDArray(arr)
	fmt.Println(twoDArray)
}

// 练习2：反转字符串
func reverse(str string) string {
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes)
}

// 练习3：一维数组转换成二维数组
func toTwoDArray(arr []int) [][]int {
	twoDArray := make([][]int, len(arr)/2)
	for i := range twoDArray {
		twoDArray[i] = make([]int, 2)
	}
	for i := 0; i < len(arr); i += 2 {
		twoDArray[i/2][0] = arr[i]
		if i+1 < len(arr) {
			twoDArray[i/2][1] = arr[i+1]
		}
	}
	return twoDArray
}