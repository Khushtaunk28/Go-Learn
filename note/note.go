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
	//struct tag =``
	Title     string    `json:""`
	Content   string    `json:""`
	CreatedAt time.Time `json:"created_time"`
}

func (note Note) Display() {
	fmt.Printf("Your Note titled %v has the following content: %v \n\n %v\n", note.Title, note.Content)
}

func (note Note) SaveToFile() error {
	fileName := strings.ReplaceAll((note.Title), " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)

}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
