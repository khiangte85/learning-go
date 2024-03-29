package main

import (
	"fmt"
	"khiangte/struct/user"
)

func main() {
	user, err := user.New("Lalmuanawma", "Khiangte", "02/6/1985")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", user)
	user.ClearUsername()
	fmt.Printf("%#v\n", user)

}
