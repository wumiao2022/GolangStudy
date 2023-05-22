package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 "Hello, World!" 到控制台
	fmt.Println("Hello, World!")

	// 练习2：声明一个整数数组，包含 1 到 10 这 10 个数字，打印数组中的所有元素
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		fmt.Println(num)
	}

	// 练习3：声明一个有 5 个元素的字符串数组，将其中的所有元素连接成一个字符串并打印出来
	strs := [5]string{"foo", "bar", "baz", "qux", "quux"}
	var result string
	for _, str := range strs {
		result += str
	}
	fmt.Println(result)

	// 练习4：声明一个 map 变量，将一些键值对存入其中，然后打印出 map 中所有的键和值
	m := map[string]string{"foo": "bar", "baz": "qux", "quux": "corge"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 练习5：声明一个切片，包含 10 个整数，使用自己喜欢的排序算法将切片中的元素排序，并打印出排序后的结果
	nums2 := []int{9, 4, 6, 2, 8, 3, 5, 7, 1, 0}
	quickSort(nums2)
	fmt.Println(nums2)
}

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	pivot := nums[0]

	var left, right []int
	for _, num := range nums[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	quickSort(left)
	quickSort(right)

	copy(nums, append(append(left, pivot), right...))
}