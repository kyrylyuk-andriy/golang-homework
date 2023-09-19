package homework4

import (
	"fmt"
	"io"
	"bufio"
)

func ReadLines(input io.Reader) ([]string) {
	var text []string

	scanner := bufio.NewScanner(input)

	InputLines := func() bool {
		fmt.Print("Add your lines here [Empty space  and Enter means stop input]: ")
		return scanner.Scan()
	}
	for InputLines() {
		eachLine := scanner.Text()
		if line == "" {
			break
		}
		text = append(text, eachLine)
	}
	
	return text
} 

func PrintLines(text []string, output io.Writer){
	for _, v := range text {
		_, err := fmt.Fprintln(output, v)
	}
}