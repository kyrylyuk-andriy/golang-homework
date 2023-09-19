package hw4

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ReadLines(input io.Reader) []string {
	var text []string

	scanner := bufio.NewScanner(input)

	InputLines := func() bool {
		fmt.Print("Add your lines here [Empty space  and Enter means stop input]: ")
		return scanner.Scan()
	}
	for InputLines() {
		eachLine := scanner.Text()
		if eachLine == "" {
			break
		}
		text = append(text, eachLine)
	}

	return text
}

func PrintLines(text []string, output io.Writer) {
	for _, v := range text {
		fmt.Fprintln(output, v)
	}
}

func SearchLines(text []string, keyword string) {
	var result []string
	for _, v := range text {
		if strings.Contains(v, keyword) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		fmt.Println("no strings found based on your search query")
	}
	fmt.Println(result)
}
