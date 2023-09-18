package main

import "fmt"

func main() {
	// 练习1：定义一个函数，对给定的整数进行判断，如果是奇数则返回true，否则返回false
	fmt.Println(isOdd(5))  // 输出：true
	fmt.Println(isOdd(10)) // 输出：false

	// 练习2：定义一个函数，输入一个字符串，统计其中的字母数量，并将每个字母以及对应的数量打印出来
	countLetters("Hello, World!") // 输出：H: 1, e: 1, l: 3, o: 2, W: 1, r: 1, d: 1

	// 练习3：定义一个函数，将给定的一组整数进行升序排序，并打印排序后的结果
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	sortNumbers(numbers) // 输出：[1 1 2 3 4 5 5 6 9]
}

func isOdd(num int) bool {
	return num%2 == 1
}

func countLetters(str string) {
	letters := make(map[rune]int)
	for _, char := range str {
		letters[char]++
	}
	for char, count := range letters {
		fmt.Printf("%c: %d, ", char, count)
	}
	fmt.Println()
}

func sortNumbers(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	fmt.Println(numbers)
}