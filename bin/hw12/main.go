package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kyrylyuk-andriy/golang-homework/hw12"
)

func main() {
	fileContent, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("error reading from file", err)
		return
	}

	inputText := string(fileContent)

	fmt.Println(inputText)

	processor := hw12.TextProcessor{}

	var action string

	flag.StringVar(&action, "action", "word-count", "Specify the action (word-count, remove-space, replace-char")

	flag.Parse()

	switch action {
	case "word-count":
		processor.SetStrategy(&hw12.WordCountStrategy{})
	case "remove-space":
		processor.AddDecorator(&hw12.RemoveSpaceDecorator{})
	case "replace-char":
		processor.AddDecorator(&hw12.ReplaceCharachterDecorator{OldChar: "a", NewChar: "A"})
	}

	result := processor.ProcessText(inputText)

	fmt.Println("processed text", result)

}
