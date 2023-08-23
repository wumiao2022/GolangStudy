package main

import "fmt"

func main() {
	// 1. 输出字符串的长度
	str := "Hello World"
	fmt.Println(len(str))

	// 2. 判断字符串是否为空
	str2 := ""
	if str2 == "" {
		fmt.Println("字符串为空")
	}

	// 3. 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 4. 查找切片中的最大值
	nums := []int{3, 5, 2, 8, 1}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	fmt.Println("最大值:", max)

	// 5. 判断一个数是否是素数
	num := 7
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Printf("%d是素数吗？%t\n", num, isPrime)
}