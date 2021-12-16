package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	horizontal := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		direction := command[0]
		value, _ := strconv.Atoi(command[1])
		if direction == "forward" {
			horizontal += value
			depth += aim*value
		} else if direction == "up" {
			aim -= value
		} else if direction == "down" {
			aim += value
		}
	}
	fmt.Print(horizontal*depth)
}