package main

import "fmt"

func main() {
	// 实例1: 打印Hello World
	fmt.Println("Hello World")
	
	// 实例2: 计算1到10之间的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 实例3: 判断一个数是否为质数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is prime")
	} else {
		fmt.Println(num, "is not prime")
	}
	
	// 实例4: 找出数组中的最大值
	nums := []int{12, 34, 56, 78, 90}
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	fmt.Println("Max:", max)
	
	// 实例5: 反转字符串
	str := "Hello"
	reversedStr := ""
	for _, c := range str {
		reversedStr = string(c) + reversedStr
	}
	fmt.Println("Reversed string:", reversedStr)
}