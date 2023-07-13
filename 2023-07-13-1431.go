package main

import "fmt"

func main() {
	// 练习1: 输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 交换两个变量的值
	a, b := 10, 20
	fmt.Printf("交换前：a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("交换后：a = %d, b = %d\n", a, b)

	// 练习3: 判断一个数是否为偶数
	num := 25
	if num%2 == 0 {
		fmt.Println("该数为偶数")
	} else {
		fmt.Println("该数为奇数")
	}

	// 练习4: 求一个数组中的最大值
	arr := []int{6, 8, 2, 10, 3}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println("最大值为：", max)

	// 练习5: 将一个字符串反转
	str := "Hello, Go!"
	reverse := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	fmt.Println("反转后的字符串：", reverse)
}