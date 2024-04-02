package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(Title, Content string) (*Note, error) {
	if Title == "" {
		return nil, errors.New("invalid Title")
	}

	if Title == "" {
		return nil, errors.New("invalid Title")
	}

	return &Note{
		Title,
		Content,
		time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Title: %v\nContent: %v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
