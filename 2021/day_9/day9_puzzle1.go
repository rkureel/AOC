package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input(f *os.File) [][]int {
	scanner := bufio.NewScanner(f)
	var numbers [][]int
	for scanner.Scan() {
		var row []int
		line := strings.Split(scanner.Text(), "")
		for i:=0;i<len(line);i++ {
			number, _ := strconv.Atoi(line[i])
			row = append(row, number)
		}
		numbers = append(numbers, row)
	}
	return numbers
}

func main() {
	f, _ := os.Open("input.txt")
	numbers := get_input(f)
	sum := 0
	for i, row := range numbers {
		for j, val := range row {
			ok := true
			if i>0 && val>=numbers[i-1][j] {
				ok = false
			}
			if i<len(numbers)-1 && val>=numbers[i+1][j] {
				ok = false
			}
			if j>0 && val>=numbers[i][j-1] {
				ok = false
			}	
			if j<len(row)-1 && val>=numbers[i][j+1] {
				ok = false
			}
			if ok {
				sum += 1+val
			}
		}
	}
	fmt.Println(sum)
}