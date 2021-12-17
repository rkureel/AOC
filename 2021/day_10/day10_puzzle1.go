package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_input(f *os.File) []string {
	scanner := bufio.NewScanner(f)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

type stack []rune

func (s *stack) empty() bool {
	if len(*s)==0 {
		return true
	}
	return false
}

func (s *stack) push(ch rune) {
	*s = append(*s, ch)
}

func (s *stack) pop() (rune, bool) {
	if s.empty() {
		return '0', true
	} else {
		element := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return element, false
	}
}

func main() {
	f, _ := os.Open("input.txt")
	input := get_input(f)
	score := 0
	bracket_map := make(map[rune]rune)
	bracket_map[')'] = '('
	bracket_map['>'] = '<'
	bracket_map[']'] = '['
	bracket_map['}'] = '{'
	for _, line := range input {
		line_score := 0
		var s stack
		for _, char := range line {
			if char=='(' || char == '[' || char == '{' || char == '<' {
				s.push(char)
			} else {
				stack_top, empty := s.pop()
				if empty || bracket_map[char] != stack_top {
					if char == ')' {
						line_score += 3
					} else if char == ']' {
						line_score += 57
					} else if char == '}' {
						line_score += 1197
					} else if char == '>' {
						line_score += 25137
					}
					break
				}
			}
		}
		score += line_score
	}
	fmt.Println(score)
}