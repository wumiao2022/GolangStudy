package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// 打印1-10的随机数
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10) + 1)
	}

	// 判断一个数是否是质数
	num := rand.Intn(100) + 1
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

// 输出示例：
// 9
// 7
// 4
// 5
// 3
// 10
// 2
// 6
// 4
// 5
// 90 is not a prime number