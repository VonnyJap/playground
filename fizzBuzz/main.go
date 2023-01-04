package main

import "fmt"

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(n int) []string {

	// from 1 to n and then check if fizz or buzz
	result := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			result = append(result, "Fizz")
			continue
		}
		if i%5 == 0 {
			result = append(result, "Buzz")
			continue
		}
		result = append(result, fmt.Sprintf("%d", i))
	}

	return result
}
