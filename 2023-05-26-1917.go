package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：打印1~100内的整数，但是如果数字是3的倍数则打印Fizz替换数字，如果数字是5的倍数则打印Buzz，如果同时是3和5的倍数则打印FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// 练习2：生成随机数组，并去除重复元素
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(20))
	}
	fmt.Println("随机数组:", arr)

	uniqueArr := []int{}
	for _, num := range arr {
		isUnique := true
		for _, uNum := range uniqueArr {
			if num == uNum {
				isUnique = false
				break
			}
		}
		if isUnique {
			uniqueArr = append(uniqueArr, num)
		}
	}
	fmt.Println("去重后数组:", uniqueArr)

	// 练习3：实现冒泡排序算法，对随机数组进行排序
	fmt.Println("排序前数组:", arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("排序后数组:", arr)
}