// https://leetcode.com/problems/valid-parentheses/description/
package main

import "fmt"

type Stack []string

func main() {
	var input string = "([])"
	boolShit := isValid(input)
	fmt.Println(boolShit)
}

func isValid(str string) bool {
	booly := true

	/* We know that, since you need a full pair of either "()", "{}" or "[]", the string cannot be true
	if the length of the string is not an even number. As such, we just return the booly variable right away.
	*/
	if (len(str) % 2) != 0 {
		booly = false
		return booly
	}

	stack := Stack{}
	for _, char := range str {
		switch char {
		case '(', '[', '{':
			stack = append(stack, string(char))
			fmt.Println(stack)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != "(" {
				fmt.Println(stack)
				booly = false
				return booly
			}
			stack = stack[:len(stack)-1]
			fmt.Println(stack)
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != "[" {
				fmt.Println(stack)
				booly = false
				return booly
			}
			stack = stack[:len(stack)-1]
			fmt.Println(stack)
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != "{" {
				fmt.Println(stack)
				booly = false
				return booly
			}
			stack = stack[:len(stack)-1]
			fmt.Println(stack)
		}
	}
	if len(stack) != 0 {
		booly = false
	}

	return booly
}

/* Not good enough ;(.
strLength := len(str)
for i := 0; i < strLength; i = i + 2 {
	if str[i] == '(' && str[i+1] != ')' {
		booly = false
		return booly
	} else if str[i] == '{' && str[i+1] != '}' {
		booly = false
		return booly
	} else if str[i] == '[' && str[i+1] != ']' {
		booly = false
		return booly
	}
}*/
