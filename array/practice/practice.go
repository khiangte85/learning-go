package main

import "fmt"

type Product struct {
	id    string
	name  string
	price float32
}

func main() {
	products := []Product{
		{
			"1", "First Product", 34.99,
		},
		{
			"2", "Second Product", 23.49,
		},
	}

	fmt.Printf("%v\n", products)
}
