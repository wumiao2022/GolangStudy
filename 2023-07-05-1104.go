package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)

	// 练习3：计算一个数组的平均值
	nums := []int{10, 20, 30, 40, 50}
	sum = 0
	for _, num := range nums {
		sum += num
	}
	average := sum / len(nums)
	fmt.Println("The average of", nums, "is", average)
  
  	// 练习4：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}
}