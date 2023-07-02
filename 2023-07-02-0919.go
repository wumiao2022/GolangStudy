package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) 

	//生成一个随机数
	num := rand.Intn(100)

	//判断随机数是奇数还是偶数
	if num%2 == 0 {
		fmt.Printf("%d是偶数\n", num)
	} else {
		fmt.Printf("%d是奇数\n", num)
	}
}