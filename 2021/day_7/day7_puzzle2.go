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

func min(a int, b int) int {
	if a>=b {
		return b
	}
	return a
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
	
	var sum float64 = 0
	for _, i := range numbers {
		sum += float64(i)
	}
	sum/=float64(len(numbers))
	h_value := int(sum)
	h_value2 := h_value+1
	cost2 := 0
	cost := 0
	for _,i := range numbers {
		val := abs(h_value-i)
		val2 := abs(h_value2-i)
		cost += val*(val+1)/2
		cost2 += val2*(val2+1)/2
	}
	fmt.Println(min(cost, cost2))
}