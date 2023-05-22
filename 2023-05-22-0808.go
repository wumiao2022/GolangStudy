package main

import (
	"fmt"
)

func main() {
	// 1. 输出"Hello, world!"
	fmt.Println("Hello, world!")
	
	// 2. 输出1到100的数字，每行输出10个
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d ", i)
		if i % 10 == 0 {
			fmt.Println()
		}
	}
	
	// 3. 从1加到100的和，并输出结果
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1+2+...+100 =", sum)
	
	// 4. 声明一个长度为5的整型数组，将数组元素赋值为1,2,3,4,5，然后按照下标从小到大的顺序输出数组元素
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	
	// 5. 声明一个map，键为字符串类型，值为整型，添加若干键值对，然后按照键从小到大的顺序输出map中的键值对
	m := map[string]int{
		"apple": 1,
		"banana": 2,
		"orange": 3,
		"pear": 4,
	}
	for k, v := range m {
		fmt.Printf("%s:%d ", k, v)
	}
}