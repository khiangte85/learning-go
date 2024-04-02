package main

import (
	"fmt"
	"khiangte/struct/user"
)

func main() {
	user1, err := user.New("Lalmuanawma", "Khiangte", "02/6/1985")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", user1)
	user1.ClearUsername()
	fmt.Printf("%#v\n", user1)

	admin := user.NewAdmin("khiangte85@gmail.com")


	fmt.Printf("%#v\n", admin)
	admin.ClearUsername()
	fmt.Printf("%#v\n", admin)

}
