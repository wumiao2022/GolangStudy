package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 计算1+2的结果并输出
	a := 1
	b := 2
	sum := a + b
	fmt.Printf("The result is: %d\n", sum)

	// 3. 判断一个数字是否为偶数并输出结果
	num := 10
	if num % 2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	// 4. 使用for循环输出1-10的数字
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 5. 定义一个切片，并添加两个元素后输出
	slice := []int{1, 2}
	slice = append(slice, 3, 4)
	fmt.Println(slice)
}