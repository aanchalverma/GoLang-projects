/*
@author : Anchal Verma
*/

package main

import "fmt"

// Function to check balanced brackets
func isbalanced(string_input string) string {
	var result = "NO"
	brackets_map := map[rune]rune{
		// opening and closing brackets mapping
		'(' : ')',
		'[' : ']',
		'{' : '}',
		// reverse mapping of closing brackets with opening brackets
		')' : '(',
		']' : '[',
		'}' : '{',

	}

	// Using stack (LIFO) operation to check balanced brackets
	stack := []rune {}

	// Looping through each character of the string
	for _, i := range string_input {
		// check if character matches with any of the closing brackets
		if i == ')' || i == '}' || i == ']' {
			// check if the bracket matches with the mapped value of the topmost element of stack
			if len(stack) > 0 && brackets_map[i] == stack[len(stack) - 1] {
				// pop the topmost element(bracket) from stack
				stack = stack[: len(stack) - 1]
				continue
			}
		}
		// else add the opening bracket to stack
		stack = append(stack, i)
	}
	// if stack is empty, all the brackets from string_input are balanced
	if len(stack) == 0 {
		result = "YES"
	}
	return result
}

// main function
func main(){

	// Taking input from user
	var string_input string
	fmt.Println("Enter the input as brackets string: ")
	fmt.Scanln(&string_input)
	
	// Print "YES" if balanced, "NO" if not balanced, or wrong input given
	fmt.Println(isbalanced(string_input))
}
