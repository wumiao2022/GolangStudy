package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印当前时间
	fmt.Println(time.Now())

	// 2. 求取3的阶乘
	factorial := 1
	for i := 1; i <= 3; i++ {
		factorial *= i
	}
	fmt.Println("3的阶乘为：", factorial)

	// 3. 判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

	// 4. 将一个切片中的元素逆序输出
	slice := []int{1, 2, 3, 4, 5}
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Print(slice[i], " ")
	}
	fmt.Println()
}