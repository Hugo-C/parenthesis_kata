package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string, error) {
	length := len(s)
	if length == 0 {
		return s, "", errors.New("Stack is empty")
	}
	return s[:length-1], s[length-1], nil
}

func IsMatchingParenthesis(open_parenthesis string, close_parenthesis string) bool {
	matchs := make(map[string]string)
	matchs["("] = ")"
	matchs["["] = "]"
	matchs["{"] = "}"

	return matchs[open_parenthesis] == close_parenthesis
}

func IsParenthesisValid(input string) bool {
	parenthesis_stack := make(stack, 0)
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == '(' || char == '[' || char == '{' {
			parenthesis_stack = parenthesis_stack.Push(string(char))
		} else if char == ')' || char == ']' || char == '}' {
			var err error
			var popped_value string
			parenthesis_stack, popped_value, err = parenthesis_stack.Pop()
			if err != nil {
				return false // Nothing to pop, meaning we have one closing parenthesis more than open ones
			}
			if !IsMatchingParenthesis(popped_value, string(char)) {
				return false // closing parenthesis doesn't match last open
			}
		}

	}
	if len(parenthesis_stack) != 0 {
		return false // we have one open parenthesis more than closed ones
	}
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
