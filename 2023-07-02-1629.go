package main

import (
	"fmt"
	"log"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个数的和并输出结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否为偶数，如果是则输出"Yes"，否则输出"No"
	number := 7
	if number%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	// 练习4: 使用for循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5: 使用for循环计算1到100的和并输出结果
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习6: 定义一个包含5个字符串元素的数组，并打印每个元素
	words := [5]string{"apple", "banana", "orange", "grape", "melon"}
	for _, word := range words {
		fmt.Println(word)
	}

	// 练习7: 定义一个切片，然后使用append函数向切片中添加元素并打印结果
	nums := []int{1, 2, 3, 4, 5}
	nums = append(nums, 6)
	fmt.Println(nums)

	// 练习8: 定义一个Map，其中键为字符串类型，值为整数类型，并打印Map中的键值对
	ages := map[string]int{
		"John":  25,
		"Alice": 30,
		"Bob":   35,
	}
	for name, age := range ages {
		fmt.Println(name, age)
	}

	// 练习9: 定义一个函数，输入两个整数，返回它们的乘积
	multiply := func(a, b int) int {
		return a * b
	}
	result := multiply(3, 4)
	fmt.Println("Result:", result)

	// 练习10: 引发一个自定义错误，并使用log包打印错误信息
	err := fmt.Errorf("This is a custom error")
	log.Println(err)
}
```