package main

import "fmt"

func main() {
	x, y := 5, 6
	num := 5

	fmt.Println(add(x, y))
	fmt.Println()
	fmt.Println(factorial(num))
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func factorial(num int) int {

	if num == 0 {

		return 1
	}
	return num * factorial(num-1)
}
