package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	message := fmt.Sprintf("Hi, %v. Welcome", name)

	if name == "" {
		return "", errors.New("name is empty")
	}

	return message, nil
}
