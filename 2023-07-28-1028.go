package main

import "fmt"

func main() {
	// 实现一个函数，将字符串中的字母按照字母表的顺序排序
	str := "golang"
	sortedStr := sortString(str)
	fmt.Println(sortedStr)
}

func sortString(str string) string {
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		for j := i + 1; j < len(bytes); j++ {
			if bytes[j] < bytes[i] {
				bytes[i], bytes[j] = bytes[j], bytes[i]
			}
		}
	}
	return string(bytes)
}