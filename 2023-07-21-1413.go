package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 随机生成一个 10 到 100 的整数
	num := rand.Intn(91) + 10

	// 判断数字是否为偶数
	if num%2 == 0 {
		fmt.Printf("%d 是一个偶数。\n", num)
	} else {
		fmt.Printf("%d 是一个奇数。\n", num)
	}

	// 判断数字是否为素数
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d 是一个素数。\n", num)
	} else {
		fmt.Printf("%d 不是一个素数。\n", num)
	}

	// 输出数字的所有因数
	fmt.Printf("%d 的所有因数有：", num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

注意：这只是一个示例，实际中可能会针对不同的练习需求编写不同的代码。