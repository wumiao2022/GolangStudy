package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：使用循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：使用条件语句判断一个数是否为偶数，并打印结果
	num := 8
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用数组存储一组整数，并使用循环打印数组中的值
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习5：使用切片追加元素，并打印切片的长度和容量
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))
}