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
	x := 0
	y := 0
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		direction := command[0]
		value, _ := strconv.Atoi(command[1])
		if direction == "forward" {
			x += value
		} else if direction == "up" {
			y += value
		} else if direction == "down" {
			y -= value
		}
	}
	fmt.Print(-1*x*y)
}