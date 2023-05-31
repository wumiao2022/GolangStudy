package main

import "fmt"

func main() {
	// 1. 定义一个数组，将数组中的元素逆序存储到另外一个数组中，并输出新数组中的元素
	arr := [5]int{1, 2, 3, 4, 5}
	reverseArr := [5]int{}
	for i := 0; i < len(arr); i++ {
		reverseArr[len(arr)-i-1] = arr[i]
	}
	fmt.Println(reverseArr)

	// 2. 定义一个map，将map中的键值对颠倒并输出
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	n := make(map[int]string)
	for k, v := range m {
		n[v] = k
	}
	fmt.Println(n)

	// 3. 实现斐波那契数列，输出前10个数
	f := [10]int{0, 1}
	for i := 2; i < 10; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	fmt.Println(f)

	// 4. 定义一个结构体，包含姓名、年龄和性别三个属性，分别使用结构体切片和map存储，并输出
	type Person struct {
		Name   string
		Age    int
		Gender string
	}
	people := []Person{
		{"Tom", 20, "Male"},
		{"Jack", 25, "Male"},
		{"Lucy", 22, "Female"},
	}
	fmt.Println(people)
	peopleMap := make(map[string]Person)
	for _, p := range people {
		peopleMap[p.Name] = p
	}
	fmt.Println(peopleMap)
}