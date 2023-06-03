package main

import (
	"fmt"
)

func main() {
	// 1. 输出Hello World
	fmt.Println("Hello World")

	// 2. 求两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("sum:", sum)

	// 3. 求三个数的平均值
	num3 := 30
	avg := (num1 + num2 + num3) / 3.0
	fmt.Println("avg:", avg)

	// 4. 判断一个数是否为偶数
	num4 := 15
	if num4%2 == 0 {
		fmt.Println(num4, "is even")
	} else {
		fmt.Println(num4, "is odd")
	}

	// 5. 计算1到100的和
	sum2 := 0
	for i := 1; i <= 100; i++ {
		sum2 += i
	}
	fmt.Println("sum2:", sum2)
	
	// 6. 将字符串反转
	str := "hello world"
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	fmt.Println("reversed:", reversed)
}