package main

import "fmt"

func main() {
	// 练习1：打印Hello World!
	fmt.Println("Hello World!")

	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 练习4：通过循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：判断一个字符串是否是回文字符串
	str := "level"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(str, "is a palindrome.")
	} else {
		fmt.Println(str, "is not a palindrome.")
	}
}