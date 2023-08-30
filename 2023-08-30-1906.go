package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算并打印1+2的结果
	sum := 1 + 2
	fmt.Println(sum)

	// 练习3：使用循环打印1-10的奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习4：判断一个数是否为偶数，如果是偶数则打印"偶数"，否则打印"奇数"
	num := 5
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习5：使用数组存储5个整数并打印
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 练习6：将字符串转换为整数并打印
	str := "123"
	num, _ := strconv.Atoi(str)
	fmt.Println(num)
}