package main

import "fmt"

// 练习1：打印Hello, World!
func main() {
	fmt.Println("Hello, World!")
}
```

```go
package main

import "fmt"

// 练习2：计算两个数的和并输出结果
func main() {
	num1 := 5
	num2 := 7
	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)
}
```

```go
package main

import "fmt"

// 练习3：判断一个数是奇数还是偶数，并输出结果
func main() {
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}
}
```

```go
package main

import "fmt"

// 练习4：计算一个整数数组的平均值并输出结果
func main() {
	nums := []int{5, 2, 7, 9, 1}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	average := float64(sum) / float64(len(nums))
	fmt.Println("The average of", nums, "is", average)
}
```

```go
package main

import "fmt"

// 练习5：判断一个字符串是否是回文，并输出结果
func main() {
	str := "racecar"
	isPalindrome := true
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(str, "is a palindrome.")
	} else {
		fmt.Println(str, "is not a palindrome.")
	}
}
```

```go
package main

import "fmt"

// 练习6：计算斐波那契数列的前n个数并输出结果
func main() {
	n := 10
	fibonacci := make([]int, n)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Println("The first", n, "numbers in the Fibonacci sequence are", fibonacci)
}
```