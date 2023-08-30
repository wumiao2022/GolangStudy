package main

import "fmt"

func main() {
	// 示例1：打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 示例2：使用两个嵌套循环打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 示例3：使用切片和循环将一个数组中的元素逆序存储
	arr := []int{1, 2, 3, 4, 5}
	reverseArr := make([]int, len(arr))
	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		reverseArr[j] = arr[i]
	}
	fmt.Println(reverseArr)
}