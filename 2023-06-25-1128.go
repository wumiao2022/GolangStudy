package main

import (
	"fmt"
)

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义一个整型变量x并赋值为10, 然后将x的值打印出来。
	x := 10
	fmt.Println(x)

	// 3. 定义两个字符串变量name和gender并分别赋值为"Tom"和"male", 然后将它们的值打印出来。
	name := "Tom"
	gender := "male"
	fmt.Println(name, gender)

	// 4. 定义一个长度为5的整型数组arr, 并将数组的元素依次赋值为1,2,3,4,5, 然后将数组打印出来。
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 5. 定义一个map结构, 其中键为"apple", 值为2.3, 键为"orange", 值为3.5, 然后将map打印出来。
	fruits := map[string]float64{
		"apple":  2.3,
		"orange": 3.5,
	}
	fmt.Println(fruits)
}