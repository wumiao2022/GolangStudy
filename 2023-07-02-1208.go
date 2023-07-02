package main

import "fmt"

func main() {
	// 1. Print the numbers from 1 to 10 using a for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2. Create a function to calculate the average of an array of float64 numbers
	numbers := []float64{1.2, 2.3, 3.4, 4.5, 5.6}
	avg := calculateAverage(numbers)
	fmt.Println("Average:", avg)

	// 3. Create a struct for a person with name and age fields
	type Person struct {
		Name string
		Age  int
	}

	// 4. Create a slice of person and print their details
	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}

	for _, person := range people {
		fmt.Println("Name:", person.Name, "Age:", person.Age)
	}

	// 5. Create a function to check if a given number is prime or not
	number := 17
	isPrime := checkPrime(number)
	fmt.Println(number, "is prime?", isPrime)
}

func calculateAverage(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	avg := total / float64(len(numbers))
	return avg
}

func checkPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}