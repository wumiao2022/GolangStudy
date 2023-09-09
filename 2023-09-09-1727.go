package main

import (
	"fmt"
	"math"
)

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello, World!")

	// 练习2：计算圆的面积和周长
	radius := 5.0
	area := math.Pi * math.Pow(radius, 2)
	circumference := 2 * math.Pi * radius
	fmt.Printf("面积：%.2f\n", area)
	fmt.Printf("周长：%.2f\n", circumference)

	// 练习3：判断一个数是否为素数
	num := 37
	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Printf("%d是素数\n", num)
	} else {
		fmt.Printf("%d不是素数\n", num)
	}

	// 练习4：找出100以内的所有素数
	for i := 2; i <= 100; i++ {
		isPrime = true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}