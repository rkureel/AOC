package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func max(a int, b int) int {
	if a>=b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a<=b {
		return a
	}
	return b
}

func get_input(f *os.File) (string, map[string]string) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	rules := make(map[string]string)
	template := scanner.Text()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		line_list := strings.Split(line, " -> ")
		rules[line_list[0]] = line_list[1]
	}
	return template, rules
}

func main() {
	f, _ := os.Open("input.txt")
	template, rules := get_input(f)
	for step:=0;step<10;step++ {
		var sb strings.Builder
		for i:=0;i<len(template)-1;i++ {
			key := template[i:i+2]
			sb.WriteByte(template[i])
			sb.WriteString(rules[key])
		}
		sb.WriteByte(template[len(template)-1])
		template = sb.String()
	}
	counts := make(map[rune]int)
	for _, char := range template {
		counts[char] += 1
	}
	max_val := 0
	min_val := len(template)
	for _, val := range counts {
		min_val = min(min_val, val)
		max_val = max(max_val, val)
	}
	fmt.Println(max_val-min_val)
}