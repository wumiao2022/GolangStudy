package main

import "fmt"

func main() {
	// 练习1：打印1到100的所有奇数
	for i := 1; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// 练习2：判断一个数是否为素数
	num := 29
	isPrime := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println(num, "是素数")
	} else {
		fmt.Println(num, "不是素数")
	}

	// 练习3：计算斐波那契数列的前20项
	n := 20
	fibonacci1, fibonacci2 := 0, 1
	fmt.Println("斐波那契数列的前20项：")
	for i := 0; i < n; i++ {
		fmt.Println(fibonacci1)
		fibonacci1, fibonacci2 = fibonacci2, fibonacci1+fibonacci2
	}

	// 练习4：统计一个字符串中每个字母出现的次数
	str := "hello world"
	counts := make(map[rune]int)
	for _, ch := range str {
		counts[ch]++
	}
	fmt.Println("每个字母出现的次数：")
	for ch, count := range counts {
		fmt.Printf("%c: %d\n", ch, count)
	}
}
