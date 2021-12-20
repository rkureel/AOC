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

var rules map[string]string
var rules_index map[string]int
var rune_index map[rune]int

func add_slices(a []int, b []int) []int {
	sum := make([]int, len(a))
	for i:=0;i<len(a);i++ {
		sum[i] = a[i]+b[i]
	}
	return sum
}

func get_input(f *os.File) (string, map[string]string) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	rules = make(map[string]string)
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
	rune_index = make(map[rune]int)
	rules_index = make(map[string]int)
	for rule, char := range rules {
		rune_index[rune(rule[0])] = 0
		rune_index[rune(rule[1])] = 0
		rune_index[rune(char[0])] = 0
		rules_index[rule] = 0
	}
	index_count := 0
	for rule := range rules_index {
		rules_index[rule] = index_count
		index_count += 1
	}
	index_count = 0
	for char := range rune_index {
		rune_index[char] = index_count
		index_count += 1
	}

	dp := make([][][]int, 40)
	for i:=0;i<40;i++ {
		dp[i] = make([][]int, len(rules))
		for j:=0;j<len(rules);j++ {
			dp[i][j] = make([]int, len(rune_index))
		}
	}
	for i:=0;i<40;i++ {
		for rule, j := range rules_index {
			if(i==0) {
				dp[i][j][rune_index[rune(rule[0])]] += 1
				dp[i][j][rune_index[rune(rules[rule][0])]] += 1
			} else {
				key_1 := rule[:1] + string(rules[rule])
				key_2 := string(rules[rule]) + rule[1:]
				val1 := dp[i-1][rules_index[key_1]]
				val2 := dp[i-1][rules_index[key_2]]
				dp[i][j] = add_slices(val1, val2)
			}
		}
	}

	counts := make([]int, len(rune_index))
	for i:=0;i<len(template)-1;i++ {
		key := template[i:i+2]
		counts = add_slices(counts, dp[39][rules_index[key]])
	}
	counts[rune_index[rune(template[len(template)-1])]] += 1
	
	max_val := counts[0]
	min_val := counts[0]
	for i:=1;i<len(counts);i++ {
		max_val = max(max_val, counts[i])
		min_val = min(min_val, counts[i])
	}
	fmt.Println(max_val-min_val)
}