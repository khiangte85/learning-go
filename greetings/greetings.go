package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	message := fmt.Sprintf(randomFormat(), name)

	if name == "" {
		return "", errors.New("name is empty")
	}

	return message, nil
}

func randomFormat() string {
	formats := []string {
		"Hi, %v, Welcome",
		"Great to see you %v",
		"Hail, %v!",
	}

	return formats[rand.Intn(len(formats))]
}
