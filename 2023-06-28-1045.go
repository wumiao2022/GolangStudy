package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印当前时间
	fmt.Println(time.Now())

	// 2. 计算1到100的和并输出
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 3. 判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "is a prime number.")
	} else {
		fmt.Println(num, "is not a prime number.")
	}

	// 4. 交换两个变量的值
	a, b := 10, 20
	fmt.Println("Before swapping, a =", a, "and b =", b)
	a, b = b, a
	fmt.Println("After swapping, a =", a, "and b =", b)
}
```