package main

import (
	"fmt"
)

func main() {
	// 1. 声明一个整数变量x，并将其初始化为10
	var x int = 10
	fmt.Println(x)

	// 2. 声明一个字符串变量name，并将其初始化为"Gopher"
	var name string = "Gopher"
	fmt.Println(name)

	// 3. 声明一个浮点数变量pi，并将其初始化为3.14
	var pi float64 = 3.14
	fmt.Println(pi)

	// 4. 声明一个布尔类型变量isTrue，并将其初始化为true
	var isTrue bool = true
	fmt.Println(isTrue)

	// 5. 声明一个数组a，包含元素1, 2, 3, 4, 5
	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	// 6. 声明一个切片b，包含元素"a", "b", "c", "d"
	var b = []string{"a", "b", "c", "d"}
	fmt.Println(b)

	// 7. 声明一个map c，包含键值对 "name": "Gopher", "age": 30, "gender": "male"
	var c = map[string]interface{}{
		"name":   "Gopher",
		"age":    30,
		"gender": "male",
	}
	fmt.Println(c)
}
