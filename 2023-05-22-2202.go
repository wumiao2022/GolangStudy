package main

import "fmt"

func main() {
	// 1. 判断一个数是否为奇数
	num1 := 5
	if num1%2 == 1 {
		fmt.Println("num1是奇数")
	}

	// 2. 求两个数的最大公约数
	num2 := 18
	num3 := 24
	for i := num2; i > 0; i-- {
		if num2%i == 0 && num3%i == 0 {
			fmt.Println(num2, "和", num3, "的最大公约数是", i)
			break
		}
	}

	// 3. 实现一个冒泡排序
	arr := []int{4, 2, 1, 6, 5, 3}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("排序后的数组：", arr)

	// 4. 打印出1-100中不能被3和5整除的数
	for i := 1; i <= 100; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
		}
	}

	// 5. 判断一个字符串是否为回文字符串
	str1 := "level"
	isPalindrome := true
	for i := 0; i < len(str1)/2; i++ {
		if str1[i] != str1[len(str1)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("字符串", str1, "是回文字符串")
	} else {
		fmt.Println("字符串", str1, "不是回文字符串")
	}
}