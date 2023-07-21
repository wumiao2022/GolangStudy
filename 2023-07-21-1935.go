package main

import "fmt"

func main() {
	// 练习1: 定义一个变量num，赋值为10，打印该变量的值
	
	num := 10
	fmt.Println(num)

	// 练习2: 声明一个整型数组arr，包含5个元素，元素的值分别为1, 2, 3, 4, 5

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 练习3: 声明一个切片slice，元素为字符串类型，包含3个元素，分别为"apple", "banana", "orange"

	slice := []string{"apple", "banana", "orange"}
	fmt.Println(slice)

	// 练习4: 使用for循环输出1到10之间的所有偶数

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 练习5: 声明一个map，包含3个键值对，键的类型为string，值的类型为int，分别为"one": 1, "two": 2, "three": 3

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)
}