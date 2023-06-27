package main

import "fmt"

func main() {
	// 1. 输出Hello, World!
	fmt.Println("Hello, World!")

	// 2. 输出斐波那契数列的前20项
	a, b := 0, 1
	for i := 0; i < 20; i++ {
		fmt.Printf("%d ", b)
		a, b = b, a+b
	}

	// 3. 输入一个年份，判断是否是闰年
	var year int
	fmt.Print("请输入一个年份：")
	fmt.Scan(&year)
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("%d年是闰年\n", year)
	} else {
		fmt.Printf("%d年不是闰年\n", year)
	}

	// 4. 将一个整数转为二进制表示的字符串
	var n int
	fmt.Print("请输入一个整数：")
	fmt.Scan(&n)
	fmt.Printf("%d的二进制表示为：%b\n", n, n)

	// 5. 对一个字符串进行反转
	s := "abcdefg"
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	fmt.Println(string(b))
}