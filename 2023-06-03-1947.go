package main

import (
	"fmt"
)

func main() {
	// 1. 输出Hello World!
	fmt.Println("Hello World!")

	// 2. 变量定义和使用
	var a int = 10
	b := 20
	fmt.Println(a + b)

	// 3. 数组定义和循环输出
	c := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}

	// 4. 切片定义和循环输出
	d := []int{6, 7, 8, 9, 10}
	for _, value := range d {
		fmt.Println(value)
	}

	// 5. Map定义和操作
	e := make(map[string]int)
	e["apple"] = 10
	e["banana"] = 5
	fmt.Println(e["apple"])
	delete(e, "banana")

	// 6. 函数定义和调用
	f := add(3, 4)
	fmt.Println(f)

	// 7. 结构体定义和使用
	type person struct {
		name string
		age  int
	}
	g := person{"Tom", 20}
	fmt.Println(g.name)

	// 8. 接口定义和使用
	type addable interface {
		add(int, int) int
	}
	type adder struct {
	}
	func (adder) add(a int, b int) int {
		return a + b
	}
	h := adder{}
	fmt.Println(h.add(3, 4))
}

func add(a int, b int) int {
	return a + b
}