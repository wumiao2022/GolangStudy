package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. 求1到100的所有数字的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100的所有数字的和为：", sum)

	// 2. 判断一个数是否为质数
	var num int
	fmt.Println("请输入一个数字：")
	fmt.Scanln(&num)
	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是质数")
	} else {
		fmt.Println(num, "不是质数")
	}

	// 3. 按从小到大的顺序输出三个数
	var a, b, c int
	fmt.Println("请依次输入三个数：")
	fmt.Scanln(&a, &b, &c)
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	fmt.Println("从小到大的排序结果为：", a, b, c)

	// 4. 判断一个字符串是否为回文字符串
	var str string
	fmt.Println("请输入一个字符串：")
	fmt.Scanln(&str)
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(str, "是回文字符串")
	} else {
		fmt.Println(str, "不是回文字符串")
	}
}