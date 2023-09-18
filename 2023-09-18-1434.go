package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数相加的结果并输出
	num1 := 10
	num2 := 5
	result := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", result)

	// 练习3：判断一个数是否是偶数并输出结果
	num := 27
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：计算一个整数数组的总和并输出
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("The sum of the numbers is:", sum)

	// 练习5：判断一个字符串是否是回文并输出结果
	str := "level"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(str, "is a palindrome")
	} else {
		fmt.Println(str, "is not a palindrome")
	}
}