package main

import (
	"fmt"
)

func main() {
	// 练习1：输入两个整数，输出它们的和
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a + b)

	// 练习2：编写程序，找出数组中最大的元素
	arr := []int{10, 20, 30, 40, 50}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)

	// 练习3：编写程序，计算字符串中每个字符出现的次数
	str := "hello world"
	countMap := make(map[rune]int)
	for _, c := range str {
		countMap[c]++
	}
	fmt.Println(countMap)

	// 练习4：编写程序，求出一个整数切片的平均值
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	avg := float64(sum) / float64(len(nums))
	fmt.Println(avg)

	// 练习5：编写函数，判断一个整数是否为素数
	fmt.Println(isPrime(7))
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}