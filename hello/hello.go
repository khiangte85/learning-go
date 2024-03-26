package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	message, err := greetings.Hello("Khiangte")

	if err != nil {
		log.SetPrefix("Greetings: ")
		log.SetFlags(0)
		log.Fatal(err)
	}

	fmt.Println(message)
}
