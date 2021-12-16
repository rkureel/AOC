package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func print_map(m map[rune]map[rune]bool) {
	for _, char := range "abcdefg" {
		fmt.Printf("\n%c->", char)
		for _, char2 := range "abcdefg" {
			fmt.Printf(" %c: %v,", char2, m[char][char2])
		}
	}
	fmt.Println()
}

func generate_number(signal_pattern []string, output []string) int {
	blueprint := make(map[int]string)
	blueprint[0] = "abcefg"
	blueprint[1] = "cf"
	blueprint[2] = "acdeg"
	blueprint[3] = "acdfg"
	blueprint[4] = "bcdf"
	blueprint[5] = "abdfg"
	blueprint[6] = "abdefg"
	blueprint[7] = "acf"
	blueprint[8] = "abcdefg"
	blueprint[9] = "abcdfg"

	reverse := make(map[string]int)
	reverse["abcefg"] = 0
	reverse["cf"] = 1
	reverse["acdeg"] = 2
	reverse["acdfg"] = 3
	reverse["bcdf"] = 4
	reverse["abdfg"] = 5
	reverse["abdefg"] = 6
	reverse["acf"] = 7
	reverse["abcdefg"] = 8
	reverse["abcdfg"] = 9
	
	char_counts := make(map[int][]int)
	char_counts[2] = []int{1}
	char_counts[3] = []int{7}
	char_counts[4] = []int{4}
	char_counts[5] = []int{2, 3, 5}
	char_counts[6] = []int{0, 6, 9}
	char_counts[7] = []int{8}
	
	m := make(map[rune]map[rune]bool)

	for _, char := range "abcdefg" {
		m[char] = make(map[rune]bool)
		for _, char2 := range "abcdefg" {
			m[char][char2] = true
		}
	}

	for _, word := range signal_pattern {
		reachable_characters := make(map[rune]bool)
		word_characters := make(map[rune]bool)
		common_characters := make(map[rune]bool)
		
		for _, char := range "abcdefg" {
			word_characters[char] = false
			reachable_characters[char] = false
			common_characters[char] = true
		}
		for _, char := range word {
			word_characters[char] = true
		}

		possible_ints := char_counts[len(word)]
		for _, integer := range possible_ints {
			for _, char := range blueprint[integer] {
				reachable_characters[char] = true
			}
			for _, char := range "abcdefg" {
				if strings.ContainsRune(blueprint[integer], char) == false {
					common_characters[char] = false
				}
			}
		}
		
		for _, char1 := range "abcdefg" {
			for _, char2 := range "abcdefg" {
				if word_characters[char1] {
					if reachable_characters[char2]==false {
						m[char1][char2] = false
					}
				} else {
					if common_characters[char2] {
						m[char1][char2] = false
					} 
				}
			}
		} 

		for _, char1 := range "abcdefg" {
			count := 0
			var temp rune
			for _, char2 := range "abcdefg" {
				if m[char1][char2] {
					count++
					temp = char2
				}
			}
			if count == 1 {
				for _, char2 := range "abcdefg" {
					if char2 != char1 {
						m[char2][temp] = false
					}
				}
			}
		}

		// fmt.Println("======================")
		// fmt.Println(word)
		// fmt.Println(possible_ints)
		// print_map(m)
	}
	// print_map(m)

	mapping := make(map[rune]rune)
	for _, char1 := range "abcdefg" {
		for _, char2 := range "abcdefg" {
			if m[char1][char2] {
				mapping[char1] = char2
			}
		}
	}

	sum := 0
	multiplier := 1000
	for _, word := range output {
		actual_string := ""
		for _, char := range word {
			actual_string += string(mapping[char])
		}

		words_list := strings.Split(actual_string, "")
		sort.Strings(words_list)
		actual_string = strings.Join(words_list, "")

		sum += reverse[actual_string]*multiplier
		multiplier/=10
		// fmt.Println(word)
		// fmt.Println(actual_string)
		// fmt.Println(sum)
	}
	return sum
}

func main() {
	f, _ := os.Open("input.txt")
	signal_patterns, outputs := get_input(f)
	// val := generate_number(signal_pattern[0], outputs[0])
	// fmt.Println(val)
	sum := 0
	for i := 0;i<len(signal_patterns);i++ {
		sum += generate_number(signal_patterns[i], outputs[i])
	}
	fmt.Println(sum)
}