package main

import "fmt"

func main() {
	// 练习1：打印整数1至10的平方数
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}

	// 练习2：判断一个数是否为偶数
	num := 8
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习3：计算一个数组的平均值
	arr := []int{5, 10, 15, 20}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	average := sum / len(arr)
	fmt.Println("平均值:", average)

	// 练习4：使用冒泡排序对一个数组进行升序排序
	arr2 := []int{4, 2, 9, 1, 6, 5}
	for i := 0; i < len(arr2)-1; i++ {
		for j := 0; j < len(arr2)-i-1; j++ {
			if arr2[j] > arr2[j+1] {
				arr2[j], arr2[j+1] = arr2[j+1], arr2[j]
			}
		}
	}
	fmt.Println("升序排序后的数组:", arr2)
}