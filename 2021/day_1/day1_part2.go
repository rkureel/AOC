package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInts(f *os.File) []int {
	scanner := bufio.NewScanner(f)
	var numbers []int
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	f, _ := os.Open("day1_part1_input.txt")
	numbers := readInts(f)
	count := 0
	for i:=0;i<len(numbers);i++ {
		if(i+3==len(numbers)) {
			break
		}
		curr := numbers[i]
		new := numbers[i+3]
		if(new>curr) {
			count +=1
		}
	}
	fmt.Print(count)
}