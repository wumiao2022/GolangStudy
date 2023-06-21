package main

import "fmt"

func main() {
	// 1. 输出Hello World!
	fmt.Println("Hello World!")

	// 2. 实现从标准输入读取一个整数N，输出1-N的所有数字
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}

	// 3. 给定一个字符串，输出其中出现频率最高的字符
	s := "hello world"
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	maxCount := 0
	maxChar := ' '
	for c, count := range m {
		if count > maxCount {
			maxCount = count
			maxChar = c
		}
	}
	fmt.Println(string(maxChar))

	// 4. 实现简单的计算器，输入两个数字和运算符（+、-、*、/），输出运算结果
	var a, b int
	var op byte
	fmt.Scanf("%d %c %d", &a, &op, &b)
	switch op {
	case '+':
		fmt.Println(a + b)
	case '-':
		fmt.Println(a - b)
	case '*':
		fmt.Println(a * b)
	case '/':
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operator!")
	}

	// 5. 实现一个自定义类型的栈，支持push、pop、peek操作
	type stack []int
	s := make(stack, 0)
	s = push(s, 1)
	s = push(s, 2)
	s = push(s, 3)
	fmt.Println(peek(s))
	s = pop(s)
	fmt.Println(peek(s))
	s = pop(s)
	fmt.Println(peek(s))
	s = pop(s)
	fmt.Println(peek(s))
}

func push(s []int, e int) []int {
	return append(s, e)
}

func pop(s []int) []int {
	if len(s) == 0 {
		return s
	}
	return s[:len(s)-1]
}

func peek(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return s[len(s)-1]
}