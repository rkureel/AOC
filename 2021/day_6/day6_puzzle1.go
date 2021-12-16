package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input(f *os.File) map[int]int {
	scanner := bufio.NewScanner(f)
	m := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ",")
		for _, number := range numbers {
			number_int, _ := strconv.Atoi(number)
			m[number_int] += 1
		}
	}
	return m
}


func main() {
	f, _ := os.Open("input.txt")
	m := get_input(f)
	var (
		i int
		j int
	)
	dp := make([][]int, 9)
	for i=0;i<9;i++ {
		dp[i] = make([]int, 81)
	}
	
	dp[0][1] = 2
	for i=1;i<9;i++ {
		dp[i][1] = 1
	}
	for i=2;i<81;i++ {
		for j=0;j<9;j++ {
			if j==0 {
				dp[j][i] = dp[6][i-1] + dp[8][i-1]
			} else {
				dp[j][i] = dp[j-1][i-1]
			}
		}
	}
	sum := 0
	for i=0;i<9;i++ {
		sum += m[i]*dp[i][80]
	}
	fmt.Println(sum)
}