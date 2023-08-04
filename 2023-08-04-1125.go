package main

import (
	"fmt"
)

func main() {
	// 1. 输出 Hello, World!
	fmt.Println("Hello, World!")

	// 2. 交换两个变量的值
	a := 5
	b := 10
	a, b = b, a
	fmt.Println("交换后的a:", a)
	fmt.Println("交换后的b:", b)

	// 3. 计算一个整数数组的和
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("数组的和为:", sum)

	// 4. 判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
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

	// 5. 将字符串反转输出
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("反转后的字符串为:", reversedStr)
}
```