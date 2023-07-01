package main

import (
	"fmt"
	"math"
)

func main() {
	// 练习1：计算圆的面积
	radius := 5.0
	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("The area of the circle is %.2f\n", area)

	// 练习2：判断一个数是否是偶数
	num := 10
	if num%2 == 0 {
		fmt.Printf("%d is an even number.\n", num)
	} else {
		fmt.Printf("%d is an odd number.\n", num)
	}

	// 练习3：计算两个数的最大公约数
	num1 := 24
	num2 := 36
	gcd := greatestCommonDivisor(num1, num2)
	fmt.Printf("The greatest common divisor of %d and %d is %d.\n", num1, num2, gcd)
}

func greatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return greatestCommonDivisor(b, a%b)
}