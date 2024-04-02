package main

import (
	"bufio"
	"errors"
	"fmt"
	"khiangte/struct-practice/note"
	"khiangte/struct-practice/todo"
	"os"
	"strings"
)

type ISaver interface {
	Save() error
}

type IOuputable interface {
	ISaver
	Display()
}

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	output(userNote)

	text, _ := getUserInput("Enter todo: ")

	todo, _ := todo.New(text)

	output(todo)
}

func output(data IOuputable) {
	data.Display()
	saveToFile(data)
}

func saveToFile(data ISaver) {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed to save data", err)
	}
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
