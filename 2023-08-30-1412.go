package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 声明整数变量a和b，并分别赋值为10和20，然后输出它们的和
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)

	// 3. 声明一个字符串变量，并赋值为"Go语言"，然后输出该字符串的长度
	str := "Go语言"
	fmt.Println("Length:", len(str))

	// 4. 声明一个切片变量slice，并进行初始化，包含元素1、2、3、4、5，然后输出该切片的长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// 5. 使用for循环打印1到10之间的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}