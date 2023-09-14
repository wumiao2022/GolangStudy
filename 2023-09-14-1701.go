package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：打印当前的时间（年、月、日、时、分、秒）
	fmt.Println(time.Now())

	// 练习2：将字符串转换为整数并输出结果
	str := "123"
	num, _ := strconv.Atoi(str)
	fmt.Println(num)

	// 练习3：将整数转换为字符串并输出结果
	num := 456
	str := strconv.Itoa(num)
	fmt.Println(str)

	// 练习4：判断一个数字是否为偶数，并输出结果
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习5：求一个数组的和，并输出结果
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Println(sum)

	// 练习6：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}

	// 练习7：使用冒泡排序对一个数组进行排序，并输出结果
	arr := []int{5, 3, 8, 2, 1}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}