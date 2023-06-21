package main

import "fmt"

func main() {
	// 打印所有1到100中能够被3整除的数
	fmt.Println("能够被3整除的数：")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	// 计算1到100的所有偶数的和
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println("1到100的所有偶数的和为：", sum)

	// 打印一个5行5列的矩阵，每个元素为当前位置的行和列的乘积
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}

	// 在控制台输入一个字符串，输出该字符串的长度、由该字符串组成的数组、该字符串的反转形式
	var s string
	fmt.Println("请输入一个字符串：")
	fmt.Scan(&s)
	fmt.Printf("该字符串的长度为：%d\n", len(s))
	fmt.Println("该字符串组成的数组为：")
	for _, c := range s {
		fmt.Printf("%c ", c)
	}
	fmt.Println()
	fmt.Println("该字符串的反转形式为：")
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()
}