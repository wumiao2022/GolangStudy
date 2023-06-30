package main

import "fmt"

func main() {
	// 1. 打印出"Hello, World!"
	fmt.Println("Hello, World!")

	// 2. 定义一个整数变量x并赋值为10，打印x的值
	x := 10
	fmt.Println(x)

	// 3. 定义一个字符串变量name并赋值为"John"，打印name的值
	name := "John"
	fmt.Println(name)

	// 4. 定义一个布尔变量isTrue并赋值为true，打印isTrue的值
	isTrue := true
	fmt.Println(isTrue)

	// 5. 定义一个整数切片numbers，并使用循环向切片中添加10个偶数，然后打印切片的内容
	numbers := make([]int, 0)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i*2)
	}
	fmt.Println(numbers)

	// 6. 定义一个结构体Person，包含name和age两个字段，并创建一个Person类型的变量p，并赋值为{name: "Alice", age: 25}，打印p的值
	type Person struct {
		name string
		age  int
	}
	p := Person{name: "Alice", age: 25}
	fmt.Println(p)

	// 7. 定义一个函数add，接收两个整数参数，并返回它们的和，调用函数并打印结果
	add := func(a, b int) int {
		return a + b
	}
	result := add(5, 3)
	fmt.Println(result)

	// 8. 定义一个具有5个元素的整数数组，使用循环将数组的元素依次加1，然后打印数组的内容
	array := [5]int{1, 2, 3, 4, 5}
	for i := range array {
		array[i] = array[i] + 1
	}
	fmt.Println(array)

	// 9. 使用if语句判断一个整数变量x的值是否大于10，如果是则打印"x大于10"，否则打印"x小于等于10"
	x = 15
	if x > 10 {
		fmt.Println("x大于10")
	} else {
		fmt.Println("x小于等于10")
	}

	// 10. 使用for循环打印输出1到10之间的所有偶数
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}