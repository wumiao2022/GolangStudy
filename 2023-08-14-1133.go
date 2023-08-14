package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 练习4：循环打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}

	// 练习5：定义一个结构体表示学生，包含姓名和年龄字段，并创建一个学生对象输出相关信息
	type Student struct {
		Name string
		Age  int
	}
	student := Student{
		Name: "Alex",
		Age:  20,
	}
	fmt.Printf("姓名：%s，年龄：%d\n", student.Name, student.Age)
}