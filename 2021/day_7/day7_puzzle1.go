package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs (x int) int {
	if x>0 {
		return x
	}
	return -x
}

func get_input(f *os.File) []int {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")
	numbers := make([]int, len(line))
	for i, n := range line {
		numbers[i], _ = strconv.Atoi(n)
	}
	sort.Ints(numbers)
	return numbers
}


func main() {
	f, _ := os.Open("input.txt")
	numbers := get_input(f)
	var n int
	n = len(numbers)
	h_value := numbers[(n+1)/2]
	sum := 0
	for _,i := range numbers {
		sum += abs(h_value-i)
	}
	fmt.Println(sum)
}