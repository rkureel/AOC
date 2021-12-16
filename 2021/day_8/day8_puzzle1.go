package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func get_input(f *os.File) ([][]string, [][]string) {
	scanner := bufio.NewScanner(f)
	var signal_patterns [][]string
	var outputs [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		signal_pattern := line[0]
		output := line[1]
		signal_patterns = append(signal_patterns, strings.Split(signal_pattern, " "))
		outputs = append(outputs, strings.Split(output, " "))
	}
	return signal_patterns, outputs
}

func main() {
	f, _ := os.Open("input.txt")
	_, outputs := get_input(f)
	unique_count := 0
	for _, output := range outputs {
		for _, word := range output {
			if len(word) == 2 || len(word) == 3 || len(word) == 4 || len(word) == 7 {
				unique_count += 1
			}
		}
	}
	fmt.Println(unique_count)
}