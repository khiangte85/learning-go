package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Text      string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(Text string) (*Todo, error) {
	if Text == "" {
		return nil, errors.New("invalid Title")
	}

	return &Todo{
		Text,
		time.Now(),
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("Text: %v\n", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
