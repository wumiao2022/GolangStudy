package main

import "fmt"

func main() {
	// 练习1：声明并初始化一个整型变量num，赋值为10，然后打印出num的值
	num := 10
	fmt.Println(num)

	// 练习2：声明并初始化两个字符串变量name1和name2，分别赋值为"John"和"Doe"，然后将两个字符串相加并打印出结果
	name1 := "John"
	name2 := "Doe"
	fullName := name1 + name2
	fmt.Println(fullName)

	// 练习3：声明一个切片变量fruits，其中包含字符串元素"apple"、"banana"和"orange"，然后打印出切片的长度和第一个元素
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println("Length of fruits:", len(fruits))
	fmt.Println("First fruit:", fruits[0])

	// 练习4：声明一个map变量scores，其中包含键值对"Math": 90, "English": 85和"Science": 95，然后打印出map的长度和"Math"对应的值
	scores := map[string]int{
		"Math":    90,
		"English": 85,
		"Science": 95,
	}
	fmt.Println("Length of scores:", len(scores))
	fmt.Println("Math score:", scores["Math"])
}