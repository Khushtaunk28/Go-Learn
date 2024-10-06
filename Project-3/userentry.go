package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"userentry/note"

	"golang.org/x/mod/sumdb/note"
)

func main() {
	title, content := GetNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	userNote.Display

}
func GetNoteData() (string, string) {
	title := GetUserInput("Note Title")
	content := GetUserInput("Note Content")
	return title, content
}

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
