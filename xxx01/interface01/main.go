package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"site.com/abc/note"
	"site.com/abc/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	printAnyKindOfValue(1)
	printAnyKindOfValue(1.5)
	printAnyKindOfValue("Hej")

	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	printAnyKindOfValue(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)

	if err != nil {
		return
	}

	userNote, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(userNote)

	if err != nil {
		return
	}
}

func printAnyKindOfValue(value interface{}) {

	intVal, ok := value.(int)
	if ok {
		fmt.Println("1a", intVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("2b", stringVal)
		return
	}

	f64Val, ok := value.(float64)
	if ok {
		fmt.Println("3c", f64Val)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println(value)
	// case string:
	// 	fmt.Println(value)
	// case float64:
	// 	fmt.Println(value)
	// default:
	// 	fmt.Println("Unknown type")
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the todo failed")
		return err
	}
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(promptText string) string {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
