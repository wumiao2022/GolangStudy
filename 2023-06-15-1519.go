package main

import "fmt"

func main() {
	// 1. 输出hello world
	fmt.Println("hello world")

	// 2. 整型变量加减乘除
	var a, b, c int
	a = 10
	b = 5
	c = a + b
	fmt.Println(c)
	c = a - b
	fmt.Println(c)
	c = a * b
	fmt.Println(c)
	c = a / b
	fmt.Println(c)

	// 3. 字符串变量拼接
	var str1, str2 string
	str1 = "Hello"
	str2 = "World"
	fmt.Println(str1 + " " + str2)

	// 4. 判断奇偶性
	var num int
	num = 10
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}

	// 5. 循环输出1-10的数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 6. 判断是否是质数
	var num2 int
	num2 = 7
	flag := true
	for i := 2; i < num2; i++ {
		if num2%i == 0 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("质数")
	} else {
		fmt.Println("非质数")
	}

	// 7. 在切片末尾添加元素
	slice1 := []int{1, 2, 3}
	slice1 = append(slice1, 4)
	fmt.Println(slice1)

	// 8. 在数组中查找元素
	arr1 := [5]int{1, 2, 3, 4, 5}
	num3 := 3
	found := false
	for _, v := range arr1 {
		if v == num3 {
			found = true
			break
		}
	}
	if found {
		fmt.Println("找到了")
	} else {
		fmt.Println("没找到")
	}

	// 9. 创建一个map并遍历它
	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	// 10. 定义一个结构体
	type Person struct {
		name string
		age  int
	}
	person1 := Person{name: "Tom", age: 18}
	fmt.Println(person1)
}