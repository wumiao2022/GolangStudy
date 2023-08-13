package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印当前时间
	fmt.Println(time.Now())

	// 2. 判断一个数是否是偶数
	num := 10
	if num % 2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 3. 切片反转
	slice := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println(slice)

	// 4. 字符串转整数
	str := "123"
	num, _ = strconv.Atoi(str)
	fmt.Println(num)

	// 5. 遍历map，打印键和值
	myMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}
	for k, v := range myMap {
		fmt.Println(k, v)
	}
}