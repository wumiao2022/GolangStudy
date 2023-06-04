package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：生成随机数
	fmt.Println(rand.Intn(100)) // 返回0-99之间的随机整数

	// 练习2：循环输出
	for i := 0; i < 5; i++ {
		fmt.Println("Hello, Golang!")
	}

	// 练习3：判断奇偶
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：数组操作
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[3])   // 返回数组中下标为3的元素
	fmt.Println(len(arr)) // 返回数组长度

	// 练习5：字符串操作
	str := "Golang is awesome!"
	fmt.Println(str[2:8])     // 返回字符串中下标从2到7的子串
	fmt.Println(len(str))    // 返回字符串长度
	fmt.Println(str + "😍") // 字符串拼接
}