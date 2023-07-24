package main

import "fmt"

func main() {
	// 1. 打印Hello, World!
	fmt.Println("Hello, World!")

	// 2. 定义一个整数变量a，值为10
	a := 10

	// 3. 定义一个字符串变量b，值为"Go语言"
	b := "Go语言"

	// 4. 打印a和b的值
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 5. 将a的值加1，并打印结果
	a++
	fmt.Println("a =", a)

	// 6. 定义一个切片c，包含元素1, 2, 3
	c := []int{1, 2, 3}

	// 7. 打印切片c的所有元素
	fmt.Println("c =", c)

	// 8. 将切片c追加一个新的元素4，并打印结果
	c = append(c, 4)
	fmt.Println("c =", c)

	// 9. 定义一个空的map，键为string类型，值为int类型
	d := make(map[string]int)

	// 10. 往map d中添加键值对，键为"apple"，值为10
	d["apple"] = 10

	// 11. 打印map d的所有键值对
	fmt.Println("d =", d)
}