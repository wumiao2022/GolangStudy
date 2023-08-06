package main

import "fmt"

func main() {
	// 1. 声明并初始化一个整型变量x，赋值为10

	x := 10
	fmt.Println(x)

	// 2. 声明一个字符串变量name，不进行初始化

	var name string
	fmt.Println(name)

	// 3. 声明一个浮点型变量pi，赋值为3.14159

	pi := 3.14159
	fmt.Println(pi)

	// 4. 声明一个布尔型变量isTrue，赋值为true

	isTrue := true
	fmt.Println(isTrue)

	// 5. 声明一个切片变量numbers，元素类型为整型，并初始化为包含1，2，3三个元素的切片

	numbers := []int{1, 2, 3}
	fmt.Println(numbers)

	// 6. 声明一个字典变量person，键为字符串类型，值为整型，并初始化为空字典

	person := make(map[string]int)
	fmt.Println(person)
}