package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note"
)

func main() {
	title, content := GetNoteData()
	userNote, err := note.New(title, content)
	userNote.Display()
	err = userNote.SaveToFile()
	if err != nil {
		fmt.Println("Saving Note failed")
		return
	}
	fmt.Println("Saving the Note succeed")
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
