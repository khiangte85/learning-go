package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	dob       string
	createdAt time.Time
}

func main() {
	// user  := User {
	// 	firstName: "Lalmuanawma",
	// 	lastName: "Khiangte",
	// 	dob: "02.06.1985",
	// 	createdAt: time.Now(),
	// }

	user1 := User{}

	fmt.Printf("%#v", user1)
}
