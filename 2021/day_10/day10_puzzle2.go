package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var scores []int
	bracket_map := make(map[rune]rune)
	bracket_score := make(map[rune]int)
	bracket_map[')'] = '('
	bracket_map['>'] = '<'
	bracket_map[']'] = '['
	bracket_map['}'] = '{'
	bracket_score['('] = 1
	bracket_score['['] = 2
	bracket_score['{'] = 3
	bracket_score['<'] = 4
	for _, line := range input {
		line_score := 0
		var s stack
		
		var corrupted bool = false
		for _, char := range line {
			if char=='(' || char == '[' || char == '{' || char == '<' {
				s.push(char)
			} else {
				stack_top, empty := s.pop()
				if empty || bracket_map[char] != stack_top {
					corrupted = true
					break
				}
			}
		}

		if corrupted {
			continue
		}

		for !s.empty() {
			char, _ := s.pop()
			line_score = line_score*5 + bracket_score[char]
		}
		scores = append(scores, line_score)
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}