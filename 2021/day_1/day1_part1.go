package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("day1_part1_input.txt")
	scanner := bufio.NewScanner(f)
	count := 0
	var (
		curr	int
		prev	int
	)
	scanner.Scan()
	prev, _ = strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		curr, _ = strconv.Atoi(scanner.Text())
		if(curr>prev) {
			count += 1
		}
		prev = curr
	}
	fmt.Println(count)
}