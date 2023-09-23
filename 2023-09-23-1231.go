package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印当前时间，格式为年-月-日 时:分:秒
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	// 2. 使用for循环输出1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3. 判断一个年份是否是闰年，并输出结果
	year := 2020
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("%d是闰年\n", year)
	} else {
		fmt.Printf("%d不是闰年\n", year)
	}

	// 4. 统计一个字符串中每个字符出现的次数，并输出结果
	str := "hello world"
	count := make(map[rune]int)
	for _, char := range str {
		count[char]++
	}
	fmt.Println(count)

	// 5. 将一个正整数转换为二进制表示，并输出结果
	num := 10
	binary := ""
	for num > 0 {
		remainder := num % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		num /= 2
	}
	fmt.Println(binary)
}