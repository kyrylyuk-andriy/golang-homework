package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("1689007676028_text.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// all words which starts from б (Cyrillic) and have exactly 3 Cyryllic  letters
	regexPattern := `\s\x{0431}\p{Cyrillic}{2}\s`
	re := regexp.MustCompile(regexPattern)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)

		for _, match := range matches {
			fmt.Println(match)
		}
	}
}
