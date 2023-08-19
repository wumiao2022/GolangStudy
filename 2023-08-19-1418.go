package main

import "fmt"

func main() {
	// 1. 判断一个数是否为偶数
	num := 10
	if num%2 == 0 {
		fmt.Println("This number is even")
	} else {
		fmt.Println("This number is odd")
	}

	// 2. 打印1到10之间的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 3. 将一个字符串反转
	str := "Hello, World!"
	reversedStr := ""
	for _, char := range str {
		reversedStr = string(char) + reversedStr
	}
	fmt.Println(reversedStr)

	// 4. 求数组中的最大值
	arr := []int{2, 5, 1, 10, 7}
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	fmt.Println("The max value is:", max)

	// 5. 计算一个正整数的平方根并保留两位小数
	num = 25
	sqrt := 0.00
	for sqrt = 1.00; sqrt*sqrt <= float64(num); sqrt += 0.01 {
	}
	sqrt = float64(int(sqrt*100)) / 100
	fmt.Println("The square root is:", sqrt)
}