package main

import "fmt"

func main() {
	// 1. 打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 声明一个整数变量x，并赋值为10
	x := 10
	fmt.Println("x =", x)

	// 3. 声明一个浮点数变量y，并赋值为3.14
	y := 3.14
	fmt.Println("y =", y)

	// 4. 声明一个字符串变量name，并赋值为"John Doe"
	name := "John Doe"
	fmt.Println("name =", name)

	// 5. 创建一个整数数组nums，包含元素1, 2, 3, 4, 5
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println("nums =", nums)

	// 6. 使用for循环打印数组nums的每个元素
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// 7. 声明一个切片slices，包含元素1, 2, 3
	slices := []int{1, 2, 3}
	fmt.Println("slices =", slices)

	// 8. 使用range关键字打印切片slices的每个元素
	for _, value := range slices {
		fmt.Println(value)
	}

	// 9. 声明一个映射m，包含键值对"a": 1, "b": 2, "c": 3
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("m =", m)

	// 10. 使用键"b"从映射m中获取值，并打印
	fmt.Println("m[\"b\"] =", m["b"])
}