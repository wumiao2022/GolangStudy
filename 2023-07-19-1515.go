package main

func main() {
	// 实例1：打印"Hello, Golang!"
	println("Hello, Golang!")

	// 实例2：两个整数相加并输出
	a := 10
	b := 5
	sum := a + b
	println(sum)

	// 实例3：循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		println(i)
	}

	// 实例4：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		println("偶数")
	} else {
		println("奇数")
	}

	// 实例5：判断一个年份是否为闰年
	year := 2020
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		println("闰年")
	} else {
		println("非闰年")
	}
}