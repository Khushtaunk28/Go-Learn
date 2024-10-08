package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type saver interface {
	Save() error
}
type Todo struct {
	Text string `json:"text"`
}

// Save implements main.saver.
func (todo Todo) Save() error {
	panic("unimplemented")
}

func (todo Todo) Display() {
	fmt.Printf("\n%v", todo.Text)
}

func (todo Todo) SaveToFile() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)

}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid Input")
	}
	return Todo{
		Text: content,
	}, nil
}
