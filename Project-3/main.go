package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note"
	"example.com/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := GetNoteData()
	todoText := GetTodoData()
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	todo.Display()
	err = saveData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)
	userNote.Display()
	err = userNote.SaveToFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Saving the Note succeed")

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("\nSaving the todo failed")
		return err
	}
	fmt.Println("\nSaving the todo succeeded")
	return nil

}

func GetTodoData() string {
	return GetUserInput("Todo Text: ")
}

func GetNoteData() (string, string) {
	title := GetUserInput("Note Title: ")
	content := GetUserInput("Note Content: ")
	return title, content
}

func GetUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
