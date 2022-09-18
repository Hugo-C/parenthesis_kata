package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsParenthesisValid(input string) bool {
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter input to validate: ")
	scanner.Scan()
	input := scanner.Text()

	res := IsParenthesisValid(input)

	formatted_input := "'" + input + "'"
	if res {
		fmt.Println(formatted_input, "is valid")
	} else {
		fmt.Println(formatted_input, "is NOT valid")
	}
}
