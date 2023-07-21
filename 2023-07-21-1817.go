package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1 - 打印当前时间
	fmt.Println(time.Now())

	// 练习2 - 声明并打印数组
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	// 练习3 - 声明并打印切片
	slice := []int{6, 7, 8, 9, 10}
	fmt.Println(slice)

	// 练习4 - 声明并打印map
	dictionary := map[string]string{"hello": "world", "go": "lang"}
	fmt.Println(dictionary)

	// 练习5 - 使用for循环打印1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习6 - 使用if条件判断打印奇偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// 练习7 - 使用switch语句打印不同的消息
	for i := 1; i <= 3; i++ {
		switch i {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		}
	}
}