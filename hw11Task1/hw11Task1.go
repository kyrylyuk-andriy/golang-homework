package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(\d{3}[-.\s]?\d{3}[-.\s]?\d{4})|\(\d{3}\)\s?\d{3}[-.\s]?\d{4}`)
	fmt.Printf("%q\n", re.FindAll([]byte(`1234567890 (123) 456-7890 (123)456-7890 123-456-7890 123.456.7890 123 456 7890`), -1))
}
