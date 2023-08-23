package main

import "fmt"

// 练习1: 输出Hello, World!
func helloWorld() {
	fmt.Println("Hello, World!")
}

// 练习2: 计算两个整数的和并输出
func sum(a, b int) {
	fmt.Println(a + b)
}

// 练习3: 判断一个数是否为偶数并输出结果
func isEven(num int) {
	if num%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("不是偶数")
	}
}

// 练习4: 将一个整数数组反转并输出
func reverseArray(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

// 练习5: 判断一个字符串是否为回文字符串并输出结果
func isPalindrome(str string) {
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("是回文字符串")
	} else {
		fmt.Println("不是回文字符串")
	}
}

func main() {
	helloWorld()
	sum(5, 7)
	isEven(10)
	reverseArray([]int{1, 2, 3, 4, 5})
	isPalindrome("level")
}
