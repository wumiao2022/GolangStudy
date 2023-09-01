package main

import "fmt"

func main() {
	// 练习1：打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习2：判断一个数是奇数还是偶数
	num := 6
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习3：计算两个数的和
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println("两个数的和为:", sum)

	// 练习4：数组遍历
	nums := []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}

	// 练习5：字符串拼接
	str1 := "Hello"
	str2 := "World!"
	result := str1 + " " + str2
	fmt.Println(result)
}