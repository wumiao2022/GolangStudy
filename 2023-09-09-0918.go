package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机数
	randomNum := rand.Intn(100)

	// 输出随机数
	fmt.Println(randomNum)

	// 定义两个数组，分别用于存放奇数和偶数
	var oddNums []int
	var evenNums []int

	// 循环生成10个随机数，并根据奇偶性存放到相应的数组中
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		} else {
			oddNums = append(oddNums, num)
		}
	}

	// 输出奇数数组
	fmt.Println("Odd numbers:", oddNums)

	// 输出偶数数组
	fmt.Println("Even numbers:", evenNums)

	// 定义一个结构体表示汽车
	type Car struct {
		Brand  string
		Color  string
		Price  int
		IsSUV  bool
		IsEV   bool
		Milage float64
	}

	// 创建一个汽车对象
	car := Car{
		Brand:  "Toyota",
		Color:  "Blue",
		Price:  30000,
		IsSUV:  false,
		IsEV:   false,
		Milage: 25000.55,
	}

	// 输出汽车对象信息
	fmt.Println(car)

	// 定义一个接口表示动物
	type Animal interface {
		Speak() string
	}

	// 定义一个结构体表示狗
	type Dog struct {
		Name  string
		Breed string
	}

	// 狗实现动物接口的Speak方法
	func (d Dog) Speak() string {
		return "Woof!"
	}

	// 创建一个狗对象
	dog := Dog{
		Name:  "Charlie",
		Breed: "Labrador",
	}

	// 输出狗对象的声音
	fmt.Println(dog.Speak())
}