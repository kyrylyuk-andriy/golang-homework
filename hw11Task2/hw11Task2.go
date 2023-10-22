package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("wiki.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// all words which starts from б and have at least 2 letters and not more than 4
	regexPattern := `\b[b]\w{1,3}\b`
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
