package main

import "fmt"

func main() {
	// 1. 声明一个整数变量x，并将其初始化为10
	x := 10

	// 2. 声明一个字符串变量name，并将其初始化为"Go语言"
	name := "Go语言"

	// 3. 打印变量x和name的值
	fmt.Println("x =", x)
	fmt.Println("name =", name)

	// 4. 声明一个浮点数变量pi，并将其初始化为3.14159
	pi := 3.14159

	// 5. 声明一个布尔变量isTrue，并将其初始化为true
	isTrue := true

	// 6. 打印变量pi和isTrue的值
	fmt.Println("pi =", pi)
	fmt.Println("isTrue =", isTrue)

	// 7. 修改变量x的值为20，并重新打印
	x = 20
	fmt.Println("x =", x)

	// 8. 声明一个新的字符串变量greeting，并将其初始化为"Hello, "
	greeting := "Hello, "

	// 9. 将字符串变量name和greeting拼接起来，并打印结果
	fmt.Println(greeting + name)
}

// 输出结果:
// x = 10
// name = Go语言
// pi = 3.14159
// isTrue = true
// x = 20
// Hello, Go语言