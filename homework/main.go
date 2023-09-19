package homework

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
		each := scanner.Text()
		if each == "" {
			break
		}
		text = append(text, each)
	}
	
	return text
} 

func PrintLines(text []string, output io.Writer){
	for _, v := range text {
		fmt.Fprintln(output, v)
	}
}
