package main

import (
	"fmt"
)

func main() {
	// 练习1：打印1到100的所有奇数
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习2：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：判断一个数是否是素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 练习4：找出数组中的最大值和最小值
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	max := nums[0]
	min := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Println("最大值：", max)
	fmt.Println("最小值：", min)

	// 练习5：翻转字符串
	str := "Hello, world!"
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))

	// 练习6：将 map 中的 key 和 value 对调
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}
	newM := make(map[int]string)
	for k, v := range m {
		newM[v] = k
	}
	fmt.Println(newM)
}