package main

import "fmt"

func main() {
	// 1. 翻转字符串
	str := "Hello, World!"
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println(reversed)

	// 2. 判断回文数字
	num := 12321
	temp := num
	reversedNum := 0
	for temp > 0 {
		reversedNum = reversedNum*10 + temp%10
		temp /= 10
	}
	isPalindrome := num == reversedNum
	fmt.Println(isPalindrome)

	// 3. 斐波那契数列
	n := 10
	fibonacci := []int{0, 1}
	for i := 2; i < n; i++ {
		fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
	}
	fmt.Println(fibonacci)

	// 4. 冒泡排序
	arr := []int{5, 3, 8, 2, 1, 4}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}