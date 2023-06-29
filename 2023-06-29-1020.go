package main

import (
	"fmt"
	"math"
)

func main() {
	// 练习1：求平方根
	num := 9.0
	sqrtNum := math.Sqrt(num)
	fmt.Printf("平方根: %f\n", sqrtNum)

	// 练习2：判断一个数是否为偶数
	num2 := 10
	if num2%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("不是偶数")
	}

	// 练习3：打印1到10之间的所有奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}

	// 练习4：计算两个数的最大公约数
	num3 := 24
	num4 := 36
	gcd := 1
	for i := 1; i <= num3 && i <= num4; i++ {
		if num3%i == 0 && num4%i == 0 {
			gcd = i
		}
	}
	fmt.Printf("最大公约数: %d\n", gcd)

	// 练习5：将字符串反转
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Printf("反转后的字符串: %s\n", reversedStr)
}