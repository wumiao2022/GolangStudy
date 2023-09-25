package main

import (
	"fmt"
)

func main() {
	// 练习1：打印星号三角形
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// 练习2：打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%2d ", j, i, i*j)
		}
		fmt.Println()
	}

	// 练习3：计算两个数的最大公约数和最小公倍数
	a, b := 36, 48

	// 计算最大公约数
	gcd := 1
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			gcd = i
		}
	}

	// 计算最小公倍数
	lcm := (a * b) / gcd

	fmt.Println("最大公约数:", gcd)
	fmt.Println("最小公倍数:", lcm)
}

// 练习4：实现一个简单的栈结构（可存放整数）
type Stack struct {
	data []int
}

func (s *Stack) Push(num int) {
	s.data = append(s.data, num)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	popValue := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return popValue
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}