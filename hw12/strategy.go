package hw12

import (
	"strconv"
	"strings"
)

type TextProcessingStrategy interface {
	ProcessText(text string) string
}

type WordCountStrategy struct{}

func (s *WordCountStrategy) ProcessText(text string) string {
	//	fmt.Println("word count strategy")
	words := strings.Fields(text)
	return strconv.Itoa(len(words))
}

type MostCommonWordsStrategy struct{}

func (s *MostCommonWordsStrategy) ProcessText(text string) string {
	return "most common word strategy"
}
