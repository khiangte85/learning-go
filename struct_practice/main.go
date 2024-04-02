package main

import (
	"bufio"
	"errors"
	"fmt"
	"khiangte/struct-practice/note"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	userNote.Save()
}

func getNoteData() (string, string) {
	title, err := getUserInput("Enter title: ")

	if err != nil {
		fmt.Println(err)
	}

	content, err := getUserInput("Enter content: ")

	if err != nil {
		fmt.Println(err)
	}

	return title, content
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		return "", errors.New("invalid input")
	}

	value = strings.ReplaceAll(value, "\n", "")

	return value, nil
}
