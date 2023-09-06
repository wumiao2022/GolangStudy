package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 两数相加
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 3. 计算两数差值
	diff := num1 - num2
	fmt.Println("Difference:", diff)

	// 4. 交换两个变量的值
	fmt.Println("Before swapping - num1:", num1, "num2:", num2)
	num1, num2 = num2, num1
	fmt.Println("After swapping - num1:", num1, "num2:", num2)

	// 5. 判断一个数是否是偶数
	checkEven(25)
	checkEven(36)

	// 6. 获取切片中的最小值
	slice := []int{5, 10, 15, 2}
	min := slice[0]
	for _, num := range slice {
		if num < min {
			min = num
		}
	}
	fmt.Println("Minimum value in the slice:", min)

	// 7. 遍历切片并打印索引和值
	for index, value := range slice {
		fmt.Println("Index:", index, "Value:", value)
	}

	// 8. 使用递归求阶乘
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Factorial of 7:", factorial(7))

	// 9. 判断一个字符串是否是回文串
	fmt.Println("Is 'racecar' a palindrome?", isPalindrome("racecar"))
	fmt.Println("Is 'hello' a palindrome?", isPalindrome("hello"))

	// 10. 定义一个结构体和方法
	rectangle := Rectangle{width: 10, height: 5}
	fmt.Println("Area of the rectangle:", rectangle.getArea())
}

func checkEven(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func isPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) getArea() int {
	return r.width * r.height
}
