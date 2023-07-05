package main

import "fmt"

func main() {
	// 1. 打印1到10之间的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 2. 根据用户输入的数字，判断是否是质数
	var num int
	fmt.Print("请输入一个数字: ")
	fmt.Scan(&num)

	isPrime := true
	for i := 2; i < num; i++ {
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

	// 3. 使用冒泡排序对一个整数切片进行排序
	numbers := []int{5, 9, 2, 8, 1, 6}
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	fmt.Println(numbers)
}