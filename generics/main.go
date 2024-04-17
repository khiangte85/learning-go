package main

import "fmt"

func main() {
	result := add(1, 3)

	fmt.Println("sum is: ", result)
}

func add[T int | float64| string](a T, b T) T {
	return a + b
}
