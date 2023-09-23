package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数
	num := rand.Intn(100)

	// 输出随机整数的平方根
	fmt.Printf("随机整数 %d 的平方根为 %f\n", num, Sqrt(float64(num)))
}

// 自定义函数计算一个浮点数的平方根
func Sqrt(x float64) float64 {
	z := x / 2 // 初始猜测值为 x 的一半

	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z) // 使用牛顿法逼近平方根
	}

	return z
}