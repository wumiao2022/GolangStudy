package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：定义一个整型变量，并打印其值
	var num int = 10
	fmt.Println(num)

	// 练习3：定义一个字符串变量，并打印其值
	var str string = "Golang is awesome!"
	fmt.Println(str)

	// 练习4：定义一个浮点型变量，并打印其值
	var f float64 = 3.14
	fmt.Println(f)

	// 练习5：定义一个布尔型变量，并打印其值
	var b bool = true
	fmt.Println(b)

	// 练习6：定义一个数组，包含5个整数，并打印每个元素的值
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 练习7：定义一个切片，并打印其长度和容量
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// 练习8：定义一个map，包含学生姓名和年龄，并打印每个键值对
	student := map[string]int{
		"Tom":   18,
		"Jerry": 20,
		"Alice": 19,
	}
	for name, age := range student {
		fmt.Println(name, age)
	}

	// 练习9：定义一个函数，接收两个整数参数，并返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(3, 5))

	// 练习10：使用条件语句判断一个数是否为正数，并打印结果
	num2 := -10
	if num2 > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Not positive")
	}
}
