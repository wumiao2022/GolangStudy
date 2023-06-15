package main

import "fmt"

func main() {
	// 1. 输出Hello World!
	fmt.Println("Hello World!")

	// 2. 声明一个变量x，输出变量x的类型和值
	var x int = 10
	fmt.Printf("x的类型是%T，值是%d\n", x, x)

	// 3. 声明一个字符串变量name并赋值，输出字符串拼接结果
	name := "Alice"
	fmt.Println("Hello " + name)

	// 4. 用循环打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 5. 声明一个切片a，保存1到5的整数，循环输出切片中的元素
	a := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()

	// 6. 声明一个map，保存学生姓名和对应的年龄，循环输出map中的键值对
	ageMap := map[string]int{"Alice": 18, "Bob": 20, "Charlie": 22}
	for name, age := range ageMap {
		fmt.Println(name, "的年龄是", age)
	}
}