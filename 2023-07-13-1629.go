package main

import "fmt"

func main() {
	// 练习1：打印Hello, world!
	fmt.Println("Hello, world!")

	// 练习2：打印整数10的二进制表示
	fmt.Printf("%b\n", 10)

	// 练习3：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习4：使用if条件判断语句判断一个数是否为偶数，并打印结果
	num := 7
	if num%2 == 0 {
		fmt.Println("是偶数")
	} else {
		fmt.Println("不是偶数")
	}

	// 练习5：定义一个切片，并在切片中添加三个元素
	slice := make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	fmt.Println(slice)

	// 练习6：使用switch语句根据条件不同打印不同的结果
	grade := 85
	switch {
	case grade >= 90:
		fmt.Println("优秀")
	case grade >= 80 && grade < 90:
		fmt.Println("良好")
	case grade >= 60 && grade < 80:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}